const htmlPdf = require('html-pdf');

class HtmlToPDF {
  constructor(html) {
    this.html = html;
  }

  async generate() {
    return new Promise((resolve, reject) => {
      htmlPdf.create(this.html).toBuffer((err, buffer) => {
        if (err) {
          reject(err);
        }
        resolve(buffer);
      });
    });
  }
}

module.exports = { HtmlToPDF };