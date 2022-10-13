const hbs = require('handlebars');

class HtmlParser {
  constructor(template) {
    this.template = template;
  }

  parse(data) {
    const template = hbs.compile(this.template);
    return template(data);
  }
}

module.exports = { HtmlParser };