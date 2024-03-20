class Library {
  constructor() {
    this.books = [];
  }

  addBook(title, author) {
    this.books.push(new Book(title, author));
  }

  checkoutBook(title, borrower) {
    this.books.forEach((book) => {
      if (book.title === title) {
        book.borrower = borrower;
        return;
      }
    });
  }

  returnBook(title) {
    this.books.forEach((book) => {
      if (book.title === title) {
        book.borrower = null;
        return;
      }
    });
  }

  listBorrowedBooks() {
    const arr = [];
    this.books.forEach((book) => {
      if (book.borrower !== null) arr.push(book);
    });
    return arr;
  }
}

class Book {
  constructor(title, author) {
    this.title = title;
    this.author = author;
    this.borrower = null;
  }
}

let library = new Library();
console.log(library);

library.addBook("The Great Gatsby", "F. Scott Fitzgerald");
library.addBook("Moby Dick", "Herman Melville");
console.log(library);

library.checkoutBook("The Great Gatsby", "John Doe");
console.log(library);

console.log("Listing borrowed books...");
console.log(library.listBorrowedBooks());

library.returnBook("The Great Gatsby");
console.log(library);
