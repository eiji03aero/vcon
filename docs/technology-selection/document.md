# Technology selection
- why do this
- general criteria
- benefits to pursue
- what to avoid by doing this
- steps

## Why do this
- Because when selecting a tool to use, we want the selected to be as most beneficial as possible to project.
- Careful selection of tool has chances to be technical debt in future.

## Benefits to pursue
- adding sufficient ability to codebase with minimal effort

## Pitfalls to avoid
- selecting unmaintained
  - unmaintained tool would probably be enough for short term, but it will likely be technical debt in future
- clutter up codebase
  - eg) paradigm mismatch, oversized tool

## General criteria
### Must ones
- features provided by the tool is enough
  - look up document to confirm features are sufficient
- how features are provided is compatible with the current codebase
  - eg) incompatible paradigm (OOP vs functional), database support
- the tool is actively maintained
  - eg) last commit within 3 month
- complexity added by the tool is acceptable
  - if it makes the codebase too complicated, probably we want to avoid it for better development experience in future

### Better to have ones
- supports unit/integrated testing
  - some tools just do not consider this, causing implementing automated testing harder
- performance is enough
  - some tools put weight in other areas like versatility, leading to performance degration
  - to analyze this, we need concrete number to achieve (eg. throughput)
- extensible in future development
  - probably hard to analyze, but better to keep in mind

## Steps
- list out candidates
  - look up by googling, articles, trend analytics (google, npm trends)
- filter candidates by if it has enough features
- filter candidates by maintenance status
- try pick couple of them by refering criteria, and make detailed comparison
