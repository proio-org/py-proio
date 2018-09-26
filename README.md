# proio for Python
[![Build Status](https://travis-ci.org/proio-org/py-proio.svg?branch=master)](https://travis-ci.org/proio-org/py-proio)
[![codecov](https://codecov.io/gh/proio-org/py-proio/branch/master/graph/badge.svg)](https://codecov.io/gh/proio-org/py-proio)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c7084aee40fd459594e3646b152d1e43)](https://www.codacy.com/app/proio-org/py-proio?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=proio-org/py-proio&amp;utm_campaign=Badge_Grade)

Please see the [main proio repository](https://github.com/proio-org/proio) for general information on proio.

## API
The API documentation is generated using Sphinx, and can be found
[here](https://proio-org.github.io/py-proio-docs/).

## Installation
Py-proio is maintained on [pypi](https://pypi.python.org/pypi/proio).  The
proio package can be installed via
`pip`:
```shell
pip install --user --upgrade proio
```

If you run into trouble importing proio, force upgrade the two dependencies
which are `protobuf` and `lz4`.

For information on what versions of Python are supported, please see the
[Travis CI page](https://travis-ci.org/proio-org/py-proio).

### Building the package
If you need to build the package (instead of just installing from pypi as
described above), you can do so with the fairly standard [setup.py](setup.py).
In this case, you need to have the Protobuf compiler (`protoc`) installed.
With `pip`:
```shell
python setup.py clean build_py
pip install --user --upgrade .
```
, or if you want to just build a source package:
```shell
python setup.py clean build_py sdist
```
, or if you want to build a wheel package:
```shell
python setup.py clean bdist_wheel
```
.  Note that the `build_py` command is responsible for generating the Protobuf
code with `protoc`, and needs to be called explicitly in the pip and source
package case.

## Examples
### Manipulating data model objects (EIC Particle)
```python
import proio
import proio.model.eic as model

test_filename = 'test_file.proio'

with proio.Writer(test_filename) as writer:
    event = proio.Event()
    
    parent = model.Particle()
    parent.pdg = 443
    parent.p.x = 1
    parent.mass = 3.097
    parent_id = event.add_entry('Particle', parent)

    child1 = model.Particle()
    child1.pdg = 11
    child1.vertex.x = 0.5
    child1.mass = 0.000511
    child1.charge = -1

    child2 = model.Particle()
    child2.pdg = -11
    child2.vertex.x = 0.5
    child2.mass = 0.000511
    child2.charge = 1

    child_ids = event.add_entries('Particle', child1, child2)
    for ID in child_ids:
        event.tag_entry(ID, 'GenStable')

    parent.child.extend(child_ids)
    child1.parent.append(parent_id)
    child2.parent.append(parent_id)

    print(event)
    writer.push(event)
```

### Iterating events in a file
```python
import proio # the model does not need to be imported when reading since proio
             # is self-descriptive!

test_filename = 'test_file.proio'
n_events = 0

with proio.Reader(test_filename) as reader:
    for event in reader:
        print('========== EVENT ' + str(n_events) + ' ==========')
        print(event)
        n_events += 1

print(n_events)
```

### Event inspection by tag
```python
import proio
import proio.model.eic as model

test_filename = 'test_file.proio'
with proio.Writer(test_filename) as writer:
    event = proio.Event()

    parent = model.Particle()
    parent.pdg = 443
    parent_id = event.add_entry('Particle', parent)

    child1 = model.Particle()
    child1.pdg = 11
    child2 = model.Particle()
    child2.pdg = -11
    child_ids = event.add_entries('Particle', child1, child2)
    for ID in child_ids:
        event.tag_entry(ID, 'GenStable')

    parent.child.extend(child_ids)
    child1.parent.append(parent_id)
    child2.parent.append(parent_id)

    writer.push(event)

with proio.Reader(test_filename) as reader:
    for event in reader:
        parts = event.tagged_entries('Particle')
        print('%i particle(s)...' % len(parts))
        for i in range(0, len(parts)):
            part = event.get_entry(parts[i])
            print('%i. PDG Code: %i' % (i, part.pdg))

            print('  %i children...' % len(part.child))
            for j in range(0, len(part.child)):
                print('  %i. PDG Code: %i' % (j, event.get_entry(part.child[j]).pdg))
```
