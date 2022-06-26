import sys
from pathlib import Path
import json

import jinja2
from markdown2 import Markdown

root_dir = Path(__file__).parent.parent
templateLoader = jinja2.FileSystemLoader(searchpath=root_dir)
env = jinja2.Environment(loader=templateLoader)

md = Markdown(html4tags=True, extras=["fenced-code-blocks", "cuddled-lists", "header-ids"])
env.filters["markdown"] = md.convert

examples = root_dir.glob("./compiler/src/test/resources/programs/example*.a60")
examples = {file.name: open(file).read() for file in examples}
example_filenames = list(examples.keys())
example_sources_json = json.dumps(examples)

result = env.get_template("website/gui-template.html").render(
    example_filenames=example_filenames, example_sources_json=example_sources_json)

with open(sys.argv[1], "w") as f:
    f.write(result)
