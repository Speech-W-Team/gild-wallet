#!/bin/bash

# Установите gomobile, если еще не установлено
go get golang.org/x/mobile/cmd/gomobile
gomobile init

# Создание библиотеки для iOS
gomobile bind -target=ios -o mobile/bindings/ios/CryptoWallet.framework ./internal/api
