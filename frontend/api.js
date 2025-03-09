// API URLs - CENTRALIZED HERE FOR EASY REPLACEMENT
const API_BASE_URL = "http://localhost:7700"; // Replace with your backend URL

const API_ENDPOINTS = {
    login: `${API_BASE_URL}/api/users/login`,
    signup: `${API_BASE_URL}/api/users/signup`,
    posts: `${API_BASE_URL}/api/posts`,
    post: (id) => `${API_BASE_URL}/api/posts?id=${id}`,
    // comments: (postId) => `${API_BASE_URL}/posts/${postId}/comments`,
    user: (id) => `${API_BASE_URL}/api/users/user?id=${id}`,
    // userPosts: (id) => `${API_BASE_URL}/users/${id}/posts`,
    // vote: (postId) => `${API_BASE_URL}/posts/${postId}/vote`
};

// API Service
const api = {
    // Auth methods
    login: async (username, password) => {
        try {
            const response = await fetch(API_ENDPOINTS.login, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ username, password })
            });
            
            if (!response.ok) {
                throw new Error('Login failed');
            }

            let data = await response.json()
            
            return {
                accessToken: data.tokenSet.accessToken,
                refreshToken: data.tokenSet.refreshToken,
                user: {
                    id: data.id,
                    username: data.username,
                    email: data.email,
                    avatar: data.avatarUrl,
                    createdAt: data.createdAt,
                    links: ["https://github.com/404th"],
                    postCount: 0
                }
            };
        } catch (error) {
            console.error('Login error:', error);
            // Fallback to mock data for development
            const mockUser = mockData.users.find(u => u.username === username);
            if (mockUser) {
                return {
                    token: "mock_jwt_token",
                    user: mockUser
                };
            }
            throw error;
        }
    }, // XXX
    
    signup: async (userData) => {
        try {
            const response = await fetch(API_ENDPOINTS.signup, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(userData)
            });
            
            if (!response.ok) {
                throw new Error('Signup failed');
            }

            let data = await response.json()
            
            return {
                accessToken: data.data.tokenSet.accessToken,
                refreshToken: data.data.tokenSet.refreshToken,
                user: {
                    id: data.data.id,
                    username: data.data.username,
                    email: data.data.email,
                    avatar: data.data.avatarUrl,
                    createdAt: data.data.createdAt,
                    links: ["https://github.com/404th"],
                    postCount: 0
                }
            }
        } catch (error) {
            console.error('Signup error:', error);
            // For development, simulate successful signup with mock data
            return {
                token: "mock_jwt_token",
                user: {
                    id: mockData.users.length + 1,
                    username: userData.username,
                    email: userData.email,
                    avatar: userData.avatarUrl || "https://via.placeholder.com/150",
                    createdAt: new Date().toISOString().split('T')[0],
                    links: [],
                    postCount: 0
                }
            };
        }
    }, // XXX
    
    // Posts methods
    getPosts: async () => {
        try {
            const response = await fetch(
                API_ENDPOINTS.posts,
                {
                    method: "GET",
                    headers: { 
                        'Content-Type': 'application/json'
                    },
                },
            );

            let data = await response.json()
            
            if (!response.ok) {
                throw new Error('Failed to fetch posts');
            }
            
            return data.posts;
        } catch (error) {
            console.error('Fetch posts error:', error);
            // Fallback to mock data
            return mockData.posts;
        }
    }, // TODO
    
    getPost: async (id) => {
        try {
            const response = await fetch(
                API_ENDPOINTS.post(id), 
                {
                    method: "GET",
                    headers: { 
                        'Content-Type': 'application/json'
                    },
                },
            );
            
            if (!response.ok) {
                throw new Error('Failed to fetch post');
            }

            let data = await response.json();

            if (data.posts.length < 1) {
                data = {posts: []}
            }
            
            return data;
        } catch (error) {
            console.error('Fetch post error:', error);
            // Fallback to mock data
            return mockData.posts.find(p => p.id === parseInt(id));
        }
    },
    
    createPost: async (postData, token) => {
        try {
            const response = await fetch(API_ENDPOINTS.posts, {
                method: 'POST',
                headers: { 
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(postData)
            });
            
            if (!response.ok) {
                throw new Error('Failed to create post');
            }
            
            let postResp = await response.json()
            
            console.log("postResp")
            console.log(postResp)
            console.log("postResp")

            return postResp;
        } catch (error) {
            console.error('Create post error:', error);
            // For development, simulate successful post creation
            const user = JSON.parse(localStorage.getItem('user'));
            return {
                id: mockData.posts.length + 1,
                title: postData.title,
                content: postData.content,
                userId: user.id,
                author: user.username,
                createdAt: new Date().toISOString().split('T')[0],
                rating: 0,
                comments: []
            };
        }
    },
    
    // Comments methods
    getComments: async (postId) => {
        try {
            const response = await fetch(API_ENDPOINTS.comments(postId));
            
            if (!response.ok) {
                throw new Error('Failed to fetch comments');
            }
            
            return await response.json();
        } catch (error) {
            console.error('Fetch comments error:', error);
            // Fallback to mock data
            return mockData.comments.filter(c => c.postId === parseInt(postId));
        }
    },
    
    createComment: async (postId, content, token) => {
        try {
            const response = await fetch(API_ENDPOINTS.comments(postId), {
                method: 'POST',
                headers: { 
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({ content })
            });
            
            if (!response.ok) {
                throw new Error('Failed to create comment');
            }
            
            return await response.json();
        } catch (error) {
            console.error('Create comment error:', error);
            // For development, simulate successful comment creation
            const user = JSON.parse(localStorage.getItem('user'));
            return {
                id: mockData.comments.length + 1,
                postId: parseInt(postId),
                userId: user.id,
                author: user.username,
                content: content,
                createdAt: new Date().toISOString().split('T')[0]
            };
        }
    },
    
    // User methods
    getUser: async (id, token) => {
        try {
            const response = await fetch(API_ENDPOINTS.user(id),
                { 
                    method: "GET",
                    headers: {
                        "Authorization": `Bearer ${token}`
                    }
                }
            );
            
            if (!response.ok) {
                throw new Error('Failed to fetch user');
            }
            
            let usr = await response.json();

            return usr
        } catch (error) {
            console.error('Fetch user error:', error);
            // Fallback to mock data
            return mockData.users.find(u => u.id === parseInt(id));
        }
    },
    
    getUserPosts: async (id) => {
        try {
            const response = await fetch(API_ENDPOINTS.userPosts(id));
            
            if (!response.ok) {
                throw new Error('Failed to fetch user posts');
            }
            
            return await response.json();
        } catch (error) {
            console.error('Fetch user posts error:', error);
            // Fallback to mock data
            return mockData.posts.filter(p => p.userId === parseInt(id));
        }
    },
    
    // Voting methods
    votePost: async (postId, vote, token) => {
        try {
            const response = await fetch(API_ENDPOINTS.vote(postId), {
                method: 'POST',
                headers: { 
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({ vote })
            });
            
            if (!response.ok) {
                throw new Error('Failed to vote on post');
            }
            
            return await response.json();
        } catch (error) {
            console.error('Vote post error:', error);
            // For development, simulate successful vote
            const post = mockData.posts.find(p => p.id === parseInt(postId));
            if (post) {
                post.rating += vote === 'up' ? 1 : -1;
                return post;
            }
            throw error;
        }
    }
};