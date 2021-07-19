export interface Book {
  title: string
  description?: string
  authors: Author[]
  type: BookType | null
  editions: { [editionName: string]: Edition }
}

export interface Author {
  name: string
}

export interface Edition {
  name: string
  date: Date
}

export enum BookType {
  hardcover = 'hardcover',
  ebook = 'ebook'
}
