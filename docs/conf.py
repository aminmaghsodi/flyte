# -*- coding: utf-8 -*-
#
# Configuration file for the Sphinx documentation builder.
#
# This file does only contain a selection of the most common options. For a
# full list see the documentation:
# http://www.sphinx-doc.org/en/stable/config

# -- Path setup --------------------------------------------------------------

# If extensions (or modules to document with autodoc) are in another directory,
# add these directories to sys.path here. If the directory is relative to the
# documentation root, use pathlib.Path.resolve(strict=True) to make it absolute, like shown here.

import os
from pathlib import Path
import logging
import sys

import sphinx.application
import sphinx.errors
from sphinx.util import logging as sphinx_logging


sys.path.insert(0, str(Path("../").resolve(strict=True)))
sys.path.append(str(Path("./_ext").resolve(strict=True)))
sys.path.append(str(Path("/Users/nikkieverett/projects/repos/flytekit")))

# import flytekit

sphinx.application.ExtensionError = sphinx.errors.ExtensionError

# -- Project information -----------------------------------------------------

project = "Flyte"
copyright = "2024, Flyte Authors"
author = "Flyte"

# The short X.Y version
version = "1.13.1"
# The full version, including alpha/beta/rc tags
release = "1.13.1"

# -- General configuration ---------------------------------------------------

# If your documentation needs a minimal Sphinx version, state it here.
#
# needs_sphinx = '1.0'

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = [
    "sphinx.ext.napoleon",
    "sphinx.ext.autodoc",
    "sphinx.ext.autosummary",
    "sphinx.ext.autosectionlabel",
    "sphinx.ext.doctest",
    "sphinx.ext.inheritance_diagram",
    "sphinx.ext.intersphinx",
    "sphinx.ext.graphviz",
    "sphinx.ext.todo",
    "sphinx.ext.coverage",
    "sphinx.ext.ifconfig",
    "sphinx.ext.viewcode",
    "sphinx.ext.extlinks",
    "sphinx-prompt",
    "sphinx_copybutton",
    "sphinx_docsearch",
    "sphinxext.remoteliteralinclude",
    "sphinx_issues",
    "sphinx_click",
    "sphinx_design",
    "sphinx_reredirects",
    "sphinxcontrib.mermaid",
    "sphinxcontrib.video",
    "sphinxcontrib.youtube",
    "sphinx_tabs.tabs",
    "sphinx_tags",
    "myst_nb",
    # custom extensions
    "auto_examples",
    "import_projects",
]

source_suffix = {
    ".rst": "restructuredtext",
    ".md": "myst-nb",
}

extlinks = {
    "propeller": ("https://github.com/flyteorg/flytepropeller/tree/master/%s", ""),
    "stdlib": ("https://github.com/flyteorg/flytestdlib/tree/master/%s", ""),
    "kit": ("https://github.com/flyteorg/flytekit/tree/master/%s", ""),
    "plugins": ("https://github.com/flyteorg/flyteplugins/tree/v0.1.4/%s", ""),
    "idl": ("https://github.com/flyteorg/flyteidl/tree/v0.14.1/%s", ""),
    "admin": ("https://github.com/flyteorg/flyteadmin/tree/master/%s", ""),
    "cookbook": ("https://flytecookbook.readthedocs.io/en/latest/", None),
}

# redirects
redirects = {
    "deprecated_integrations/index": "../../flytesnacks/deprecated_integrations/index.html",
    "deprecated_integrations/mmcloud_plugin/index": "../../flytesnacks/examples/mmcloud_agent/index.html",
    "deprecated_integrations/mmcloud_plugin/mmcloud_plugin_example": "../../flytesnacks/examples/mmcloud_agent/index.html",
    "deprecated_integrations/bigquery_plugin/index": "../../flytesnacks/deprecated_integrations/bigquery_plugin/index.html",
    "deprecated_integrations/bigquery_plugin/bigquery_plugin_example": "../../flytesnacks/deprecated_integrations/bigquery_plugin/bigquery_plugin_example.html",
    "deprecated_integrations/databricks_plugin/index": "../../flytesnacks/deprecated_integrations/databricks_plugin/index.html",
    "deprecated_integrations/databricks_plugin/databricks_plugin_example": "../../flytesnacks/deprecated_integrations/databricks_plugin/databricks_plugin_example.html",
    "deprecated_integrations/snowflake_plugin/index": "../../flytesnacks/deprecated_integrations/snowflake_plugin/index.html",
    "deprecated_integrations/snowflake_plugin/snowflake_plugin_example": "../../flytesnacks/deprecated_integrations/snowflake_plugin/snowflake_plugin_example.html",
}

autosummary_generate = True
suppress_warnings = ["autosectionlabel.*", "myst.header"]
autodoc_typehints = "description"

# The master toctree document.
master_doc = "index"

# The language for content autogenerated by Sphinx. Refer to documentation
# for a list of supported languages.
#
# This is also used if you do content translation via gettext catalogs.
# Usually you set "language" from the command line for these cases.
language = "en"

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This pattern also affects html_static_path and html_extra_path .
exclude_patterns = [
    "_build",
    "Thumbs.db",
    ".DS_Store",
    "auto/**/*.ipynb",
    "auto/**/*.py",
    "auto/**/*.md",
    "examples/**/*.ipynb",
    "examples/**/*.py",
    "jupyter_execute/**",
    "README.md",
    "_projects/**",
    "_src/**",
    "examples/**",
    "flytesnacks/index.md",
    "flytesnacks/README.md",
    "flytekit/**/README.md",
    "flytekit/_templates/**",
    "flytekit/**/*.rst",
    "flytekit/*.md",
    "flytekit/**/*ipynb",
    "flytekit/**/*.md",
    "api/flyteidl/boilerplate/**",
    "api/flyteidl/tmp/**",
    "api/flyteidl/gen/**",
    "api/flyteidl/README.md",
    "flytesnacks/sg_execution_times.rst"
]

# -- Options for HTML output -------------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_favicon = "images/favicon-flyte-docs.png"
html_logo = "images/favicon-flyte-docs.png"
html_theme = "pydata_sphinx_theme"
html_title = "Flyte"

templates_path = ["_templates"]

pygments_style = "tango"
pygments_dark_style = "native"

html_theme_options = {
    # custom flyteorg pydata theme options
    "github_url": "https://github.com/flyteorg/flyte",
    "icon_links": [
        {
            "name": "GitHub",
            "url": "https://github.com/flyteorg/flyte",
        }
    ]
}

# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".
# html_static_path = ["_static"]

# Custom sidebar templates, must be a dictionary that maps document names
# to template names.
#
# The default sidebars (for documents that don't match any pattern) are
# defined by theme itself.  Builtin themes are using these templates by
# default: ``['localtoc.html', 'relations.html', 'sourcelink.html',
# 'searchbox.html']``.

html_sidebars = {
    "api/**": ["sidebars/custom"],
    "cluster_deployment/**": ["sidebars/custom"],
    "community/**": ["sidebars/custom"],
    "concepts/**": ["sidebars/custom"],
    "ecosystem/**": ["sidebars/custom"],
    "flytesnacks/integrations/**": ["sidebars/integrations"],
    "flytesnacks/tutorials/**": ["sidebars/tutorials"],
    "user_guide/**": ["sidebars/custom"]
}

html_context = {
    "dir_to_title": {
        "api": "API",
        "cluster_deployment": "Cluster deployment",
        "community": "Community",
        "concepts": "Flyte concepts",
        "ecosystem": "Ecosystem",
        "integrations": "Integrations",
        "tutorials": "Tutorials",
        "user_guide": "User guide"
    }
}

# -- Options for HTMLHelp output ---------------------------------------------

# Output file base name for HTML help builder.
htmlhelp_basename = "Flytedoc"

# -- Options for LaTeX output ------------------------------------------------

latex_elements = {
    # The paper size ('letterpaper' or 'a4paper').
    #
    # 'papersize': 'letterpaper',
    # The font size ('10pt', '11pt' or '12pt').
    #
    # 'pointsize': '10pt',
    # Additional stuff for the LaTeX preamble.
    #
    # 'preamble': '',
    # Latex figure (float) alignment
    #
    # 'figure_align': 'htbp',
}

# Grouping the document tree into LaTeX files. List of tuples
# (source start file, target name, title,
#  author, documentclass [howto, manual, or own class]).
latex_documents = [
    (master_doc,
     "Flyte.tex",
     "Flyte Documentation",
     "Flyte Authors",
     "manual"),
]

# -- Options for manual page output ------------------------------------------

# One entry per manual page. List of tuples
# (source start file, name, description, authors, manual section).
man_pages = [(master_doc, "flyte", "Flyte Documentation", [author], 1)]

# -- Options for Texinfo output ----------------------------------------------

# Grouping the document tree into Texinfo files. List of tuples
# (source start file, target name, title, author,
#  dir menu entry, description, category)
texinfo_documents = [
    (
        master_doc,
        "Flyte",
        "Flyte Documentation",
        author,
        "Flyte",
        "Accelerate your ML and data workflows to production.",
        "Miscellaneous",
    ),
]

# Algolia docsearch credentials
docsearch_app_id = "WLG0MZB58Q"
docsearch_api_key = "28bf9bfd4a77a7d6b3ab7e98c671e781"
docsearch_index_name = "flyte"

# -- Extension configuration -------------------------------------------------
# autosectionlabel_prefix_document = True
autosectionlabel_maxdepth = 2

# Tags config
tags_create_tags = True
tags_extension = ["md", "rst"]
tags_page_title = "Tag"
tags_overview_title = "Pages by tags"

# -- Options for intersphinx extension ---------------------------------------

# Example configuration for intersphinx: refer to the Python standard library.
# intersphinx configuration. Uncomment the repeats with the local paths and update your username
# to help with local development.
intersphinx_mapping = {
    "python": ("https://docs.python.org/3", None),
    "numpy": ("https://numpy.org/doc/stable", None),
    "pandas": ("https://pandas.pydata.org/pandas-docs/stable/", None),
    "torch": ("https://pytorch.org/docs/master/", None),
    "scipy": ("https://docs.scipy.org/doc/scipy/reference", None),
    "matplotlib": ("https://matplotlib.org", None),
    "pandera": ("https://pandera.readthedocs.io/en/stable/", None),
}

myst_enable_extensions = ["colon_fence"]
myst_heading_anchors = 6

# Sphinx-mermaid config
mermaid_output_format = "raw"
mermaid_version = "latest"
mermaid_init_js = "mermaid.initialize({startOnLoad:true});"

# Makes it so that only the command is copied, not the output
copybutton_prompt_text = "$ "

# prevent css style tags from being copied by the copy button
copybutton_exclude = 'style[type="text/css"]'

nb_execution_mode = "off"
nb_execution_excludepatterns = [
    "api/flytekit/**/*",
    "flytesnacks/**/*",
    "examples/**/*",
]
nb_custom_formats = {
    ".md": ["jupytext.reads", {"fmt": "md:myst"}],
}

# These patterns are used to replace values in source files that are imported
# from other repos.
REPLACE_PATTERNS = {
    r"</auto_examples": r"</flytesnacks/examples",
    r"/auto_examples": r"/flytesnacks/examples",
    r"<flytesnacks/examples": r"</flytesnacks/examples"
}

import_projects_config = {
    "clone_dir": "_projects",
    "source_regex_mapping": REPLACE_PATTERNS,
    "dev_build": bool(int(os.environ.get("MONODOCS_DEV_BUILD", 1))),
}

# Define these environment variables to use local copies of the projects. This
# is useful for building the docs in the CI/CD of the corresponding repos.
flytesnacks_local_path = os.environ.get("FLYTESNACKS_LOCAL_PATH", None)
flytekit_local_path = os.environ.get("FLYTEKIT_LOCAL_PATH", None)

flytesnacks_path = flytesnacks_local_path or "_projects/flytesnacks"
flytekit_path = flytekit_local_path or "_projects/api/flytekit"

import_projects = [
    {
        "name": "flytesnacks",
        "source": flytesnacks_local_path or "https://github.com/flyteorg/flytesnacks",
        "docs_path": "docs",
        "dest": "flytesnacks",
        "cmd": [
            ["cp", "-R", f"{flytesnacks_path}/examples", "./examples"],
            [
                # remove un-needed docs files in flytesnacks
                "rm",
                "-rf",
                "flytesnacks/jupyter_execute",
                "flytesnacks/auto_examples",
                "flytesnacks/_build",
                "flytesnacks/_tags",
                "flytesnacks/index.md",
                "examples/advanced_composition",
                "examples/basics",
                "examples/customizing_dependencies",
                "examples/data_types_and_io",
                "examples/development_lifecycle",
                "examples/extending",
                "examples/productionizing",
                "examples/testing"
            ]
        ],
        "local": flytesnacks_local_path is not None,
    },
    {
        "name": "flytectl",
        "source": "../flytectl",
        "docs_path": "docs/source",
        "dest": "api/flytectl",
        "local": True,
    },
    {
        "name": "flyteidl",
        "source": "../flyteidl",
        "docs_path": "protos",
        "dest": "api/flyteidl",
        "cmd": ["cp", "../flyteidl/README.md", "api/flyteidl/README.md"],
        "local": True,
    }
]

# myst notebook docs customization
auto_examples_dir_root = "examples"
auto_examples_dir_dest = "flytesnacks/examples"
auto_examples_refresh = False

# -- Options for todo extension ----------------------------------------------

# If true, `todo` and `todoList` produce output, else they produce nothing.
todo_include_todos = True

# -- Options for Sphinx issues extension --------------------------------------

# GitHub repo
issues_github_path = "flyteorg/flyte"

# equivalent to
issues_uri = "https://github.com/flyteorg/flyte/issues/{issue}"
issues_pr_uri = "https://github.com/flyteorg/flyte/pull/{pr}"
issues_commit_uri = "https://github.com/flyteorg/flyte/commit/{commit}"


# Disable warnings from flytekit
os.environ["FLYTE_SDK_LOGGING_LEVEL_ROOT"] = "50"

# Disable warnings from tensorflow
os.environ["TF_CPP_MIN_LOG_LEVEL"] = "3"


class CustomWarningSuppressor(logging.Filter):
    """Filter logs by `suppress_warnings`."""

    def __init__(self, app: sphinx.application.Sphinx) -> None:
        self.app = app
        super().__init__()

    def filter(self, record: logging.LogRecord) -> bool:
        msg = record.getMessage()

        # TODO: These are all warnings that should be fixed as follow-ups to the
        # monodocs build project.
        filter_out = (
            "duplicate label",
            "Unexpected indentation",
            'Error with CSV data in "csv-table" directive',
            "Definition list ends without a blank line",
            "autodoc: failed to import module 'awssagemaker' from module 'flytekitplugins'",
            "Enumerated list ends without a blank line",
            'Unknown directive type "toc".',  # need to fix flytesnacks/contribute.md
        )

        if msg.strip().startswith(filter_out):
            return False

        if (
            msg.strip().startswith("document isn't included in any toctree")
            and record.location == "_tags/tagsindex"
        ):
            # ignore this warning, since we don't want the side nav to be
            # cluttered with the tags index page.
            return False

        return True


def setup(app: sphinx.application.Sphinx) -> None:
    """Setup root logger for Sphinx"""
    logger = logging.getLogger("sphinx")

    warning_handler, *_ = [
        h for h in logger.handlers
        if isinstance(h, sphinx_logging.WarningStreamHandler)
    ]
    warning_handler.filters.insert(0, CustomWarningSuppressor(app))
