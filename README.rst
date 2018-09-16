=======
GoError
=======

A go interface for easily handeling error calls.

Credit
------

This package uses ANSI Escape Codes from `github.com/fatih/color <https://github.com/fatih/color>`_
by `Fatih Arslan <https://github.com/fatih>`_.

Installation
------------

.. code:: go

   import (
      "github.com/TheSp1der/GoError"
      "errors"
   )

Usage
-----

Fatal Errors: This will stop the process.

.. code:: go

   GoError.Fatal(errors.New("This is a fatal error and will call os.Exit()"))

Warning Errors: Identify the error as a warning.

.. code:: go

   GoError.Warning(errors.New("This is a warning"))

Informational Errors:

.. code:: go

   GoError.Info(errors.New("This is non-critical information"))


