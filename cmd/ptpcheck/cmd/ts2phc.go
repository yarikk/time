/*
Copyright (c) Facebook, Inc. and its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/facebook/time/phc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	srcDeviceTS2PHCFlag string
	dstDeviceTS2PHCFlag string
	intervalTS2PHCFlag  time.Duration
	firstStepTS2PHCFlag time.Duration
	srcPinTS2PHCFlag    uint
	dstPinTS2PHCFlag    uint
)

func init() {
	RootCmd.AddCommand(ts2phcCmd)
	ts2phcCmd.Flags().StringVarP(&srcDeviceTS2PHCFlag, "source", "s", "/dev/ptp_tcard", "Source for Time of Day (ToD) data. Only PHC devices are supported at the moment")
	ts2phcCmd.Flags().StringVarP(&dstDeviceTS2PHCFlag, "destination", "d", "eth0", "PHC to be synchronized. The clock may be identified by its character device (like /dev/ptp0) or its associated network interface (like eth0).")
	ts2phcCmd.Flags().DurationVarP(&intervalTS2PHCFlag, "interval", "i", time.Second, "Interval between syncs in nanosseconds")
	ts2phcCmd.Flags().DurationVarP(&firstStepTS2PHCFlag, "first_step", "f", 20*time.Microsecond, "The maximum offset, specified in seconds, that the servo will correct by changing the clock frequency instead of stepping the clock. This is only applied on the first update. When set to 0.0, the servo will not step the clock on start. The default is 0.00002 (20 microseconds).")
	ts2phcCmd.Flags().UintVarP(&srcPinTS2PHCFlag, "out-pin", "n", phc.DefaultTs2PhcIndex, "output pin number of the PPS signal on source device. Defaults to Pin 3")
	ts2phcCmd.Flags().UintVarP(&dstPinTS2PHCFlag, "in-pin", "o", phc.DefaultTs2PhcSinkIndex, "input pin number of the PPS signal on destination device. Defaults to Pin 3")
}

func ts2phcRun(srcDevicePath string, dstDeviceName string, interval time.Duration, stepth time.Duration, srcPinIndex uint) error {
	ppsSource, err := getPPSSourceFromPath(srcDevicePath, srcPinIndex)
	if err != nil {
		return fmt.Errorf("error opening source phc device: %w", err)
	}
	dstDevice, err := phcDeviceFromName(dstDeviceName)
	if err != nil {
		return fmt.Errorf("error opening target phc device: %w", err)
	}
	ppsSink, err := phc.PPSSinkFromDevice(dstDevice, dstPinTS2PHCFlag)
	if err != nil {
		return fmt.Errorf("error setting target device as PPS sink: %w", err)
	}
	pi, err := phc.NewPiServo(interval, stepth, dstDevice)
	if err != nil {
		return fmt.Errorf("error getting servo: %w", err)
	}
	ticker := time.NewTicker(interval)

	for range ticker.C {
		eventTime, err := phc.PollLatestPPSEvent(ppsSink)
		if err != nil {
			log.Errorf("Error polling PPS Sink: %v", err)
			continue
		} else if eventTime.IsZero() {
			continue
		}
		log.Debugf("PPS event at %+v", eventTime.UnixNano())
		if err := phc.PPSClockSync(pi, ppsSource, eventTime, dstDevice); err != nil {
			log.Errorf("Error syncing PHC: %v", err)
		}
	}
	return nil
}

func getPPSSourceFromPath(srcDevicePath string, pinIndex uint) (*phc.PPSSource, error) {
	srcDeviceFile, err := os.Open(srcDevicePath)
	if err != nil {
		return nil, fmt.Errorf("error opening source device: %w", err)
	}
	ppsSource, err := phc.ActivatePPSSource(phc.FromFile(srcDeviceFile), pinIndex)
	if err != nil {
		return nil, fmt.Errorf("error activating PPS Source for PHC: %w", err)
	}

	return ppsSource, nil
}

// phcDeviceFromName returns a PHC device from a device name, which can be either an interface name or a PHC device path
func phcDeviceFromName(dstDeviceName string) (*phc.Device, error) {
	devicePath, err := phc.IfaceToPHCDevice(dstDeviceName)
	if err != nil {
		log.Infof("Provided device name is not an interface, assuming it is a PHC device path")
		devicePath = dstDeviceName
	}
	// need RW permissions to issue CLOCK_ADJTIME on the device
	dstFile, err := os.OpenFile(devicePath, os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("error opening destination device: %w", err)
	}
	dstDevice := phc.FromFile(dstFile)
	return dstDevice, nil
}

var ts2phcCmd = &cobra.Command{
	Use:   "ts2phc",
	Short: "Sync PHC with external timestamps",
	Run: func(_ *cobra.Command, _ []string) {
		ConfigureVerbosity()
		if err := ts2phcRun(srcDeviceTS2PHCFlag, dstDeviceTS2PHCFlag, intervalTS2PHCFlag, firstStepTS2PHCFlag, srcPinTS2PHCFlag); err != nil {
			log.Fatal(err)
		}
	},
}