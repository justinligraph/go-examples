package main

var index_html = `
<html>
  <body>
    Hello world from <p/>
    %s
    <form action="/" method="post">
      <input type="text" name="name"/> <p/>
      <input type="submit"/>
    </form>
  </body>
</html>
`
