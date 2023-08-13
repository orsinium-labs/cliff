# cliff

The simplest and safest golang library for making CLI tools.

## Simplicity

...

## Safety

The following is checked at compilation time:

* You don't use values before parsing arguments.
* You don't initialize the same flag twice.
* You use only one letter for flag shorthands.
* You provide default values of the correct types.
* You pass pointers, not values, to flag targets.
* You use only types supported by the library.
