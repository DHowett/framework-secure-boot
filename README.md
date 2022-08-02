## Framework Laptop UEFI Secure Boot Certificates

**Source**: Extracted from live machines.

**Date**: 2021-10-21 (TGL), 2022-08-02 (ADL)

Certificates marked with `ADL` are for the 12th generation Intel Framework Laptop.

_With thanks to @ngxson for providing the ADL certificates._

### KEK (Key Exchange Key)

This certificate allows Framework to add additional certificates to the DB.

Current thumbprints:

- `frame.work-LaptopKEK.der`: `a1c36bd73e143c77954be234f71337370136074f`
- `frame.work-LaptopADLKEK.der`: `55d6ef7949b5b18c5bb7895605c1c4de18943e94`

### DB (Database)

This certificate allows for the automatic trust of any Framework-signed EFI binaries (such as the
BIOS updater.)

Current thumbprints:

- `frame.work-LaptopDB.der`: `732bcb5921f51141a8cd6ff213e4aad43cbb6adc`
- `frame.work-LaptopADLDB.der`: `a6519d24d7f31c548f5f42c2e736a41be7ce622f`

### Notes

The Framework certificates, when in EFI signature list format, have an owner GUID of
`55555555-5555-5555-5555-555555555555`. The `.esl` files in this repository bear the
same owner GUID.

#### Verification

> **Note**
> If you are on a 12th generation Intel Framework Laptop, substitute `ADLDB` and `ADLKEK` in the
> following commands.

You have no reason to trust me, a stranger on the internet, to provide certificates that
can verify boot applications on your machine.

What you can trust, however, is the default key set on your own laptop. Use the following
commands to extract the default `PK` (Platform Key) and validate the `KEK`.

_Assuming efivarfs is mounted at `/sys/firmware/efi/efivars`_

```bash
# Read PKDefault (skip the 4-byte EFI variable flags)
dd if=/sys/firmware/efi/efivars/PKDefault-8be4df61-93ca-11d2-aa0d-00e098032b8c of=pk.bin skip=4 iflag=skip_bytes

# Convert EFI signature list to certificates (package efitools on Debian, Ubuntu)
sig-list-to-certs pk.bin pk

# Convert the PK certificate pk-0.der to PEM; do the same for the included KEK
openssl x509 -inform der -outform pem -in pk-0.der -out pk-0.crt
openssl x509 -inform der -outform pem -in frame.work-LaptopKEK.der -out frame.work-LaptopKEK.crt

# Validate KEK was issued by the PK you just extracted
openssl verify -trusted ./pk-0.crt ./frame.work-LaptopKEK.crt
```

You should hope to see `./frame.work-LaptopKEK.crt: OK`.

Once you have established suitable trust in the `KEK`, you can use it to validate the `DB`.

**NOTE**: This is complicated by the fact that the `DB` certificate was issued by the `KEK`, but
the `KEK` *does not permit certificate signing*. OpenSSL will complain about this and fail to
validate the chain.

Therefore, a small Go program is provided in `verify.go` to verify that one certificate was issued
by another without validating that the issuing certificate is, in fact, a CA. It's only around 20
lines, and if you're going to use it to trust these certificates you may want to look at the code.

```bash
go run verify.go frame.work-LaptopDB.der frame.work-LaptopKEK.der
```

Once again, look for `frame.work-LaptopDB.der: OK`.
