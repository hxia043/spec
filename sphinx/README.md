**Q: What's the Sphinx?**  
A: Sphinx is the doc generator/framework which can generate the doc with specific format(pdf/html) according to the specific input(reStructuredText or Markdown). Use Sphnix can simplify the edit, deployment and maintain of doc or html cause the function provided by Sphinx.   

**Q: How to use the Sphinx?**  
A:
1) install the Sphinx:
```
$ pip3 install -U Sphinx
```

2) To quickly start, can run `sphinx-quickstart` to see a demo of Sphinx as: 
```
$ sphinx-quickstart
```

3) With the demo sphinx struct:
```
sphinx
|---- build
|---- source
|----- _static
|----- _templates
|----- conf.py
|----- index.rst
|- make.bat
|- Makefile
```

- conf.py is the conf info of Sphinx doc.
- index.rst is the entry of the Sphinx doc, the index can bind multiple docs with one index/directory/link.
- _static and _templates is the sub directory of the doc, it can defined the rst to describe the sub index.

For build directory, with `sphinx-build` can generate the different type of doc like:
```
$ sphinx-build -b html sourcedir builddir
```

or run make to build the output by Makefile like:
```
$ make html
```

For check the website of doc, can `Open With the Live Server` in build/html/index.html by vs code tool.
After finished can deployment the build into nginx server which provide the service of web.

**Q: reference of Sphinx?**   
A:   
- [quickstart](https://www.sphinx-doc.org/zh_CN/master/usage/quickstart.html)
- [reStructuredText](https://www.sphinx-doc.org/zh_CN/master/usage/restructuredtext/directives.html#)
