# Introduction

Please respect the following rules:

- Remain polite, courteous, respectful and constructive
throughout the evaluation process. The well-being of the community
depends on it.
- Identify with the person (or the group) evaluated the eventual
dysfunctions of the work. Take the time to discuss
and debate the problems you have identified.
- You must consider that there might be some difference in how your
peers might have understood the project's instructions and the
scope of its functionalities. Always keep an open mind and grade
him/her as honestly as possible. The pedagogy is valid only and
only if peer-evaluation is conducted seriously.
Guidelines
- Only grade the work that is in the student or group's
GiT repository.
- Double-check that the Git repository belongs to the student
or the group. Ensure that the work is for the relevant project
and also check that "git clone" is used in an empty folder.
- Check carefully that no malicious aliases was used to fool you
and make you evaluate something other than the content of the
official repository.
- To avoid any surprises, carefully check that both the evaluating
and the evaluated students have reviewed the possible scripts used
to facilitate the grading.
- If the evaluating student has not completed that particular
project yet, it is mandatory for this student to read the
entire subject prior to starting the defence.
- Use the flags available on this scale to signal an empty repository,
non-functioning program, a norm error, cheating etc. In these cases,
the grading is over and the final grade is 0 (or -42 in case of
cheating). However, with the exception of cheating, you are
encouraged to continue to discuss your work (even if you have not
finished it) in order to identify any issues that may have caused
this failure and avoid repeating the same mistake in the future.
- Remember that for the duration of the defence, no segfault,
no other unexpected, premature, uncontrolled or unexpected
termination of the program, else the final grade is 0. Use the
appropriate flag.

You should never have to edit any file except the configuration file if it exists.
If you want to edit a file, take the time to explicit the reasons with the
evaluated student and make sure both of you are okay with this.

You are allowed to use any of the different tools available on the computer, such as
pprof.

## Mandatory Part

Randomaliens

Project structure
- Check that project only requires running `go build` to produce an executable
- Check that submission includes *go.mod*, *go.sum* and *\*.proto*
- Yes No

### Task 00

- Check that gRPC server code can be generated from provided *.proto* file
Yes No
- Check that gRPC entries have 'session_id' as a string, 'frequency' as a double and also a current timestamp in UTC
Yes No
- Check that generated value for mean always falls into [-10, 10] interval
Yes No
- Check that generated value for standard deviation always falls into [0.3, 1.5] interval
Yes No
- Check that generated values for session ID/expected value/standard deviation are written into log by server
Yes No
- Check that session_id, expected value and standard deviation are generated uniquely for every new incoming client connection
Yes No
- Check that frequency values are sampled from a normal distribution with generated mean/STD
Yes No

### Task 01

- Check that client accepts '-k' parameter (STD coefficient) as a float value
Yes No
- Check that entries aren't piling up in a process memory, so there is no leak in the long run
Yes No
- Check that estimated values for expected value/standard deviation are close to ones that were used by the server
Yes No
- Check that anomalies (values which differ from mean by more than "k * STD") are detected and logged correctly
Yes No

### Task 02

- Check that database entries have 'session_id' as a string, 'frequency' as a double and also a current timestamp in UTC
Yes No
- Check that ORM is used for writing objects into PostgreSQL
Yes No
- Check that all detected anomalies are being correctly written into a database
Yes No

### Task 03

- Check that the whole pipeline (server generates a stream of entries -> client receives them -> client predicts distribution parameters -> client detect anomalies -> client writes anomalies into PostgreSQL) works
Yes No

## Ratings

Don't forget to check the flag corresponding to the defense

Ok
Outstanding project
Empty work
Invalid compilation
Crash
Leaks
Wrong result

## Conclusion

Leave a comment on this evaluation