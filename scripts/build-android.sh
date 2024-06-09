#!/bin/bash

# Установите gomobile, если еще не установлено
go get golang.org/x/mobile/cmd/gomobile
gomobile init

# Создание библиотеки для Android
gomobile bind -target=android -o mobile/bindings/android/CryptoWallet.aar ./internal/api
