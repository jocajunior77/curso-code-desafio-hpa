package main

import "testing"

func TestSum(t *testing.T ) {
    msg := greeting()

    if msg != "Code.education Rocks!" {
        t.Errorf("Saudacao invalida")
    }

}