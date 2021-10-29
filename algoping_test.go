package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSandboxCode(t *testing.T) {
	
	result, time, mem := Sandbox(lang, filePath, inPath, outPath) 
	assert.NotNil(t, result, "No result was found")
	assert.NotNil(t, time, "No elapsed time was found")
	assert.NotNil(t, mem, "No memory usage was found")
}



