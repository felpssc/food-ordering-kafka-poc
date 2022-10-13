const fs = require('fs').promises;
const path = require('path');

const { HtmlParser } = require('./HtmlParser');


class Invoice {
  constructor(order) {
    this.order = order;
  }

  async generate() {
    const template = await fs.readFile(path.resolve(__dirname, '../templates/invoice.hbs'), 'utf8');
    const parser = new HtmlParser(template);
    return parser.parse(this.order);
  }
}

module.exports = { Invoice };