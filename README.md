# lilbro
A userspace tracker for everything you do on your computer while you code.

# Why lilbro
As found by [Maalej et al](https://mobis.informatik.uni-hamburg.de/wp-content/uploads/2014/06/TOSEM-Maalej-Comprehension-PrePrint2.pdf) (page 25):
>...we asked developers whether they would share knowledge if their development environments would automatically capture it. ... Overall, respondents are rather willing to share knowledge automatically. More than 85% of respondents would share the context of an encountered problem and actions performed to solve it. About 75% agreed to share other knowledge such as the goal of changing a particular piece of code or sources consulted while solving a particular problem. This kind of knowledge would facilitate program comprehension tasks.

`lilbro` takes this idea to the simplest and extreme manifestation - a tracker for everything that you could do on a computer so that you can review it later, reflect on it and potentially share it.

![Screenshot](https://raw.githubusercontent.com/groktools/lilbro/master/lilbro.png)

# Design
`lilbro` has three logical components:

* A **server** that exposes an HTTP endpoint `/track` that allows you to send it data, which it dutifully stores in a log file.
* Numerous **client**s that send the server data. Clients could be:
  * IDE plugins
  * Browser plugins
  * OS-based keyboard and mouse hooks
  * Anything that you'd want to capture.
* **Aggregators** that make sense out of all this collected data. The design of this component is TBD.

To keep things above board, it is not a stealth tracker but an overt, userspace one. That is, you have to explicitly start and stop it.

## Data format

`lilbro` saves data in an easily processed open format:It is:
* line-based format making it amenable to grep-style processing.
* CSV format to make it amenable to processing using Excel-type tools.
* designed with a fixed set of initial fields like timestamp, context, origin app and session ID
* designed to allow custom fields using a `name=value,name=value` format within the CSV one.
* designed to allow structured data with a `data="{...json...}"` field. If present, this must be the last field.

# Portability
`lilbro` is built in `Go` and released as linux, OSX and Windows executables for easy adoption and use.

# Status
3/6/2016: `lilbro` can capture tracker information and display it in the console, it doesnt yet store the data.
