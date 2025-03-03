// Mock Data for when API is not available
const mockData = {
    users: [
        {
            id: 1,
            username: "johndoe",
            email: "john@example.com",
            avatar: "./assets/whiteLineBlackBackgroundUserIcon.png",
            createdAt: "2024-11-15",
            links: [
                { title: "Twitter", url: "https://twitter.com" },
                { title: "GitHub", url: "https://github.com" }
            ],
            postCount: 5
        },
        {
            id: 2,
            username: "janedoe",
            email: "jane@example.com",
            avatar: "./assets/whiteLineBlackBackgroundUserIcon.png",
            createdAt: "2024-12-01",
            links: [
                { title: "Portfolio", url: "https://portfolio.com" }
            ],
            postCount: 3
        },
        {
            id: 3,
            username: "bobsmith",
            email: "bob@example.com",
            avatar: "./assets/whiteLineBlackBackgroundUserIcon.png",
            createdAt: "2025-01-10",
            links: [],
            postCount: 7
        }
    ],
    
    posts: [
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 1,
            title: "Getting Started with Go Programming",
            content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It's a great language for beginners and experienced developers alike. In this post, I'll share my journey learning Go and some tips to help you get started.\n\nOne of the things I love about Go is its simplicity. The syntax is clean and easy to understand, making it accessible for newcomers. The standard library is comprehensive, providing everything you need to build web servers, work with files, and handle common programming tasks.\n\nGo's built-in concurrency features are another highlight. Goroutines and channels make it straightforward to write programs that make efficient use of multiple cores, without getting bogged down in complex threading code.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-15",
            rating: 42,
            comments: [1, 2]
        },
        {
            id: 2,
            title: "Building RESTful APIs with Go",
            content: "In this post, I'll walk through creating a RESTful API using Go and some popular packages. We'll cover setting up routes, connecting to a database, and implementing authentication.\n\nGo is an excellent choice for building APIs due to its performance and simplicity. The language's standard library already provides most of what you need to build HTTP servers, and the ecosystem offers many mature packages for the rest.\n\nWe'll be using the Gin framework for routing, as it offers a good balance of performance and features. For database access, GORM provides a clean interface for working with various databases.",
            userId: 1,
            author: "johndoe",
            createdAt: "2025-02-20",
            rating: 35,
            comments: [3]
        },
        {
            id: 3,
            title: "My Thoughts on Frontend Designed in 2025",
            content: "The frontend landscape continues to evolve rapidly. In this post, I share my perspectives on current trends and where things might be heading.\n\nOne trend I've noticed is the continued rise of component-based architectures. Frameworks like React, Vue, and Svelte have made it easier than ever to build applications from reusable components. This approach has improved developer productivity and code maintainability.\n\nAnother interesting development is the growing popularity of utility-first CSS frameworks like Tailwind. By providing a set of low-level utility classes, these frameworks allow for rapid UI development without the overhead of writing custom CSS.",
            userId: 2,
            author: "janedoe",
            createdAt: "2025-02-25",
            rating: 28,
            comments: []
        },
        {
            id: 4,
            title: "Implementing Authentication in Go Applications",
            content: "Authentication is a critical component of most applications. In this post, I'll explain how to implement JWT-based authentication in a Go web application.\n\nJSON Web Tokens (JWTs) provide a stateless way to handle user authentication. Instead of storing session information on the server, the server issues a signed token that contains the user's identity and possibly some permissions. The client then includes this token in subsequent requests to authenticate itself.\n\nIn Go, there are several packages that can help with JWT implementation. We'll be using the popular jwt-go package, which provides a clean API for working with tokens.",
            userId: 3,
            author: "bobsmith",
            createdAt: "2025-03-01",
            rating: 39,
            comments: [4, 5]
        },
        {
            id: 5,
            title: "Optimizing Database Queries in Go",
            content: "Database performance can make or break your application. In this post, I'll share some techniques for optimizing database queries in Go applications.\n\nOne of the most important aspects of database optimization is proper indexing. By adding indexes to columns that are frequently used in WHERE clauses, ORDER BY statements, or joins, you can dramatically improve query performance.\n\nAnother important consideration is the N+1 query problem, which occurs when you fetch a collection of items and then execute a query for each item to get related data. This can be addressed by using eager loading or join queries. Database performance can make or break your application. In this post, I'll share some techniques for optimizing database queries in Go applications. One of the most important aspects of database optimization is proper indexing. By adding indexes to columns that are frequently used in WHERE clauses, ORDER BY statements, or joins, you can dramatically improve query performance. Another important consideration is the N+1 query problem, which occurs when you fetch a collection of items and then execute a query for each item to get related data. This can be addressed by using eager loading or join queries.",
            userId: 3,
            author: "bobsmith",
            createdAt: "2025-03-02",
            rating: 31,
            comments: []
        }
    ],
    
    comments: [
        {
            id: 1,
            postId: 1,
            userId: 2,
            author: "janedoe",
            content: "Great introduction to Go! I've been thinking about learning it for a while, and this post has convinced me to give it a try.",
            createdAt: "2025-02-16"
        },
        {
            id: 2,
            postId: 1,
            userId: 3,
            author: "bobsmith",
            content: "I've been using Go for a few years now, and I completely agree with your points. The simplicity and performance make it a joy to work with.",
            createdAt: "2025-02-17"
        },
        {
            id: 3,
            postId: 2,
            userId: 3,
            author: "bobsmith",
            content: "Gin is definitely a great choice for Go APIs. I've used it in several projects and it's been very reliable.",
            createdAt: "2025-02-21"
        },
        {
            id: 4,
            postId: 4,
            userId: 1,
            author: "johndoe",
            content: "Very clear explanation of JWT authentication. I'll definitely be referring back to this post for my next project.",
            createdAt: "2025-03-01"
        },
        {
            id: 5,
            postId: 4,
            userId: 2,
            author: "janedoe",
            content: "I've been struggling with implementing proper authentication in my Go apps. This post is exactly what I needed!",
            createdAt: "2025-03-02"
        }
    ]
};