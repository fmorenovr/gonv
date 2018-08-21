package gonv

import(
  "errors";
)

// list of errors
var(
  GonvErrNotFound     = errors.New("Environmental Variable not found.\n")
  GonvErrAlreadyExist = errors.New("Environmental Variable is already exist.\n")
)
