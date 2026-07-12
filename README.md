# KyraFS

> **Secure Storage Processing Engine**

KyraFS (KRYptography + stoRAge File Service) is a lightweight storage engine designed for **secure file processing, storage, and streaming**. Unlike traditional object storage systems, KyraFS focuses on providing a simple, portable, and secure file processing engine that can be integrated into any application through a minimal HTTP API.

---

# Philosophy

KyraFS is built around a few simple principles:

* **Simple**
* **Fast**
* **Portable**
* **Secure**
* **Framework Independent**

KyraFS is **not** a CMS, web framework, or application platform.

KyraFS is a **storage processing engine**.

---

# Features

* Secure file storage
* Zero temporary decrypted files
* Automatic image conversion (WebP)
* Automatic video conversion (MP4)
* File inspection & MIME detection
* Cryptographic encryption (libsodium)
* Cross-platform support
* Lightweight HTTP daemon
* Framework independent
* Volume management
* Zero database dependency
* Stream-first architecture

---

# Architecture

```text
               Client
                  │
                  │ HTTP
                  ▼
          KyraFS Daemon (Go)
                  │
                  ▼
        KyraFS Engine (Python)
                  │
      ┌───────────┴───────────┐
      │                       │
  Crypto Engine         File Processor
      │                       │
      └───────────┬───────────┘
                  │
               Storage
```

---

# Core Principles

* No MVC
* No ORM
* No Framework
* No Database
* No Temporary Decrypted Files
* Streaming First
* Cross Platform
* Single Responsibility

---

# Storage Layout

```text
storage/

A10/
    BF3/
        coffee.jpg/
            8fd92ab7.syn
```

The original filename is preserved as the folder name, while all metadata and file content are stored inside the `.syn` container.

---

# Secure Mode

When Secure Mode is enabled, uploaded files are processed through the following pipeline:

```text
Original File
      │
      ▼
Inspector
      │
      ▼
Converter
      │
      ▼
Encryption
      │
      ▼
Storage (.syn)
```

When the file is requested:

```text
.syn
   │
   ▼
Decrypt Stream
   │
   ▼
Browser
```

No decrypted file is ever written to disk.

---

# Public Mode

In Public Mode, files are stored without encryption and retain their original format.

---

# HTTP API

### Native URL

```http
GET /{hash}/{filename}
```

Example:

```http
GET /A10Bf3n5jk7hko/coffee.jpg
```

---

### Clean URL

```http
GET /coffee.jpg

X-KyraFS-Hash: A10Bf3n5jk7hko
```

---

### Secure Access

```http
X-KyraFS-Key: your-secure-key
```

If a secure file is requested without a valid security key:

**401 Unauthorized**

If the provided key is invalid:

**403 Forbidden**

---

# Technology Stack

### Daemon

* Go

### Processing Engine

* Python

### Image Processing

* Pillow

### Video Processing

* FFmpeg

### Cryptography

* libsodium (PyNaCl)

---

# Roadmap

## v0.1

* File format
* Metadata engine
* Encryption
* Storage engine

## v0.2

* Streaming engine
* HTTP daemon
* Command-line interface

## v0.3

Official SDKs

* PHP
* Python
* Node.js

## v1.0

Stable Release

---

# Design Goals

* Minimal dependencies
* Easy deployment
* High performance
* Secure by default
* Framework agnostic
* Clean architecture
* Easy integration
* Production ready

---

# Project Status

🚧 **Active Development**

KyraFS is currently under active development and is **not yet recommended for production environments**.

---

# License

The project license will be finalized before the first public release.

---

# Author

KyraFS is an independent open-source project focused on secure file processing, storage, and streaming.

