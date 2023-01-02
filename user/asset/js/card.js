// jQuery
$('form').card({
  container: '.card-wrapper'
});
 
// Vanilla JavaScript
new Card({
  form: document.querySelector('form'),
  container: '.card-wrapper'
});
// Class bilgilerine dikkat!

$('form').card({
 
  // numara formatı
  formatting: true,
 
  // form seçicileri
  formSelectors: {
    numberInput: 'input[name="number"]',
    expiryInput: 'input[name="expiry"]',
    cvcInput: 'input[name="cvc"]',
    nameInput: 'input[name="name"]'
  },
  cardSelectors: {
    cardContainer: '.jp-card-container',
    card: '.jp-card',
    numberDisplay: '.jp-card-number',
    expiryDisplay: '.jp-card-expiry',
    cvcDisplay: '.jp-card-cvc',
    nameDisplay: '.jp-card-name'
  },
 
  // mesajlar
  messages: {
    validDate: 'valid\nthru',
    monthYear: 'month/year'
  },
 
  // placeholder alanları
  placeholders: {
    number: '&bull;&bull;&bull;&bull; &bull;&bull;&bull;&bull; &bull;&bull;&bull;&bull; &bull;&bull;&bull;&bull;',
    cvc: '&bull;&bull;&bull;',
    expiry: '&bull;&bull;/&bull;&bull;',
    name: 'Full Name'
  },
 
  // input masking aktifle
  masks: {
    cardNumber: false
  },
 
  // valid/invalid CSS class bilgileri
  classes: {
    valid: 'jp-card-valid',
    invalid: 'jp-card-invalid'
  },
 
  // debug mode
  debug: false
  
});