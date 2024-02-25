[![codecov](https://codecov.io/github/ilario-pierbattista/parallel-phpunit-executor/graph/badge.svg?token=HWZJAdwnhX)](https://codecov.io/github/ilario-pierbattista/parallel-phpunit-executor)

# What is that thing for

Provide a dramatically fast and memory safe way to execute phpunit tests in parallel.

# Desidered workflow

- given a folder and a batch size, split the test files in chunks. 
    - each chunk must have a unique identifier
- find (or request as input), the phpunit config file
    - duplicate and encode chunks as test suites
- for a given parallelism N, execute N workers to work on the chunks X. Each worker will tackle a chunk at time:
    - execute phpunit on the chunk
    - gather junit report and cobertura coverage report
- merge results
    - merge junit files
    - merge cobertura files
- create other artifacts
    - create html and txt report from merged cobertura

## Other notes

- as soon as a worker ends with a chunk, it should pass to another one
- output interleaving: the simples way is to wait for a worker to finish to print its output
- maybe there is a way to directly interface with PHP world and phpunit internals, but it's way to early to think about that
