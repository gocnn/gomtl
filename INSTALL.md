# GOMTL Installation Guide

A minimalist guide for installing GOMTL, a Go library for Metal-based GPU computing on macOS.

> **Quick Start**: Ensure your Mac has Metal-capable hardware before proceeding. See [Apple's Metal documentation](https://developer.apple.com/metal/) for details.

## System Requirements

- **OS**: macOS 10.15 (Catalina) or later
- **Hardware**: Metal-capable Mac (Apple Silicon or supported Intel Mac)
- **Xcode Command Line Tools**: Required for compilation

## Prerequisites

1. **Install Xcode Command Line Tools**:

   ```bash
   xcode-select --install
   ```

   Follow the prompts to complete installation.

2. **Verify Metal Support**:

   - Navigate to **Apple menu > About This Mac > More Info > System Report > Graphics/Displays**.
   - Confirm "Metal: Supported" is listed for your graphics hardware.

## Installation

1. Install GOMTL:

   ```bash
   go install github.com/gocnn/gomtl/cmd/gomtl@latest
   ```

2. Test GOMTL:

   ```bash
   gomtl
   ```

## Verification

Run the following command to verify installation:

```bash
gomtl
```

If successful, you should see output similar to:

```bash
          _____                   _______                   _____                _____                    _____  
         /\    \                 /::\    \                 /\    \              /\    \                  /\    \ 
        /::\    \               /::::\    \               /::\____\            /::\    \                /::\____\
       /::::\    \             /::::::\    \             /::::|   |            \:::\    \              /:::/    /
      /::::::\    \           /::::::::\    \           /:::::|   |             \:::\    \            /:::/    / 
     /:::/\:::\    \         /:::/~~\:::\    \         /::::::|   |              \:::\    \          /:::/    /  
    /:::/  \:::\    \       /:::/    \:::\    \       /:::/|::|   |               \:::\    \        /:::/    /   
   /:::/    \:::\    \     /:::/    / \:::\    \     /:::/ |::|   |               /::::\    \      /:::/    /    
  /:::/    / \:::\    \   /:::/____/   \:::\____\   /:::/  |::|___|______        /::::::\    \    /:::/    /     
 /:::/    /   \:::\ ___\ |:::|    |     |:::|    | /:::/   |::::::::\    \      /:::/\:::\    \  /:::/    /      
/:::/____/  ___\:::|    ||:::|____|     |:::|    |/:::/    |:::::::::\____\    /:::/  \:::\____\/:::/____/       
\:::\    \ /\  /:::|____| \:::\    \   /:::/    / \::/    / ~~~~~/:::/    /   /:::/    \::/    /\:::\    \       
 \:::\    /::\ \::/    /   \:::\    \ /:::/    /   \/____/      /:::/    /   /:::/    / \/____/  \:::\    \      
  \:::\   \:::\ \/____/     \:::\    /:::/    /                /:::/    /   /:::/    /            \:::\    \     
   \:::\   \:::\____\        \:::\__/:::/    /                /:::/    /   /:::/    /              \:::\    \    
    \:::\  /:::/    /         \::::::::/    /                /:::/    /    \::/    /                \:::\    \   
     \:::\/:::/    /           \::::::/    /                /:::/    /      \/____/                  \:::\    \  
      \::::::/    /             \::::/    /                /:::/    /                                 \:::\    \ 
       \::::/    /               \::/____/                /:::/    /                                   \:::\____\
        \::/____/                 ~~                      \::/    /                                     \::/    /
                                                           \/____/                                       \/____/
Detected 1 Metal Capable device(s)
Device 0: "Apple M4"
  General Device Information:
    Low-power:                               No
    Removable:                               No
    Configured as headless:                  No
    Registry ID:                             4294968168
  GPU Family Support:
    Apple family 1 GPU features:             Yes
    Apple family 2 GPU features:             Yes
    Apple family 3 GPU features:             Yes
    Apple family 4 GPU features:             Yes
    Apple family 5 GPU features:             Yes
    Apple family 6 GPU features:             Yes
    Apple family 7 GPU features:             Yes
    Apple family 8 GPU features:             Yes
    Mac family 2 GPU features:               Yes
    Common family 1 GPU features:            Yes
    Common family 2 GPU features:            Yes
    Common family 3 GPU features:            Yes
    Metal 3 features:                        Yes
Result = PASS
```

This output confirms:

- Metal is properly detected
- GOMTL can successfully interface with Metal drivers
- Your GPU specifications are correctly identified
- All device attributes are accessible

If you see this output, GOMTL is installed and working correctly. You can now proceed to use GOMTL in your projects.
