//go:build !V7
// +build !V7

package main

func preventExtraArgs(args []string) error {
	return nil
}
