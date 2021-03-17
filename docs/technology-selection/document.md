# Technology selection
- why do this
- benefits to pursue
- pitfalls to avoid
- general criteria
- steps

## Why do this
- Because when selecting a tool to use, we want the selected to be as most beneficial as possible to project.
- Careless selection of tool will likely be technical debt in future.

## Benefits to pursue
- adding sufficient ability to codebase with minimal effort
- improve user experience
- contribute to sales
- improve development productivity
- logically explain the reason behind why tool was chosen, so that it can be as persuasive as possible to any of colleague

## Pitfalls to avoid
- selecting unmaintained tool
  - unmaintained tool would probably be enough for short term, but it will likely be technical debt in future
- clutter up codebase
  - eg) paradigm mismatch, oversized tool

## Conditions to clarify
- service requirements
  - eg) security, functionality, 
- financial conditions
  - how much cost we can spend
- development resource
  - number of people
  - development experience

## General criteria
### Must ones
- features provided by the tool is enough
  - look up document to confirm features are sufficient
- how features are provided is compatible with the current codebase
  - eg) incompatible paradigm (OOP vs functional), database support
- the tool is actively maintained
  - eg) last commit within 3 month
- complexity added by the tool is acceptable
  - if it makes the codebase too complicated, probably we want to avoid it for mainteinability

### Better to have ones
- supports unit/integrated testing
  - some tools just do not consider this, causing implementing automated testing harder
- performance is enough
  - some tools put weight in other areas like versatility, leading to performance degration
  - to analyze this, we need concrete number to achieve (eg. throughput)
- extensible in future development
  - probably hard to analyze, but better to keep in mind

## Steps
- clarify conditions
- list out the features wanted
- list out candidates
  - look up by googling, articles, trend analytics (google, npm trends)
- filter candidates by maintenance status
- filter candidates by if it has enough features
- try pick couple of them by refering criteria, and make detailed comparison

## References
- https://note.com/timakin/n/n2d9eaa02e633
  - 4 criterias:
    - meets service requirement
    - financial status
    - development productivity
    - technological reliability
- https://qiita.com/daijinload/items/cab3ed862c08da4c06e8
- https://diary.shift-js.info/tech-and-org-and-security/
  - clarify assumptions
    - what is the expected responsibility for the code to write?
      - should care about hardware spec?
      - should care about os functionality/capability?
  - decision will inevitably have impact on:
    - operation, extending codebase
  - have to build and share criteria
    - so that we can make conclusion logically
- https://logmi.jp/tech/articles/322647
