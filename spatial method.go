package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	testutil "github.com/Pylons-tech/pylons/testutil/cli"
)

const badRecipeLiteral = `
{
    "cookbookID": "cookbookLoudTest",
    "ID": "LOUDGetCharacter",
    "name": "LOUD-Get-Character-Recipe",
    "description": "Creates a basic character in LOUD (but don't b/c it doesn't work)",
    "version": "v0.0.1",
    "coinInputs": [],
	"beef": "edible",
    "itemInputs": [],
    "entries": {
}}
