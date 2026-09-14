
lambda { |stdout,stderr,status|
  output = stdout + stderr
  # go names the package after the module, so the reason it could not run is
  # the bracket it puts after that name rather than the path. [build failed]
  # is the compiler refusing a source file, [setup failed] is go/parser
  # refusing a _test.go file before the compiler ever sees it.
  return :amber if /\[build failed\]/.match(output)
  return :amber if /\[setup failed\]/.match(output)

  # go exits 1 for a test that failed and 2 for a test binary that died, so
  # a panic is told apart from a disagreement by the status go prints.
  return :amber if /^exit status 2$/.match(output)

  return :amber if /syntax error:/.match(output)
  return :red   if /FAIL/.match(output)

  # go test prints PASS for a binary that ran no tests at all, saying so only
  # in this warning. Without it a learner who comments out their only test
  # gets a green light.
  return :amber if /no tests to run/.match(output)

  return :green if /PASS/.match(output)
  return :amber
}
