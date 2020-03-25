package main

import "testing"

func TestSum(t *testing.T ) {
    msg := greeting("Code.education Rocks!")

    if msg != "Code.education Rocks!" {
        t.Errorf("Saudacao invalida")
    }

}