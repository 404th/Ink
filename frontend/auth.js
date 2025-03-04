// Auth Service for handling user authentication
const auth = {
    // Initialize auth state
    init: () => {
        const token = localStorage.getItem('token');
        const user = localStorage.getItem('user');
        
        // If token and user exist, consider user logged in
        if (token && user) {
            return { isAuthenticated: true, token, user: JSON.parse(user) };
        }
        
        return { isAuthenticated: false, token: null, user: null };
    },
    
    // Login user
    // TODO
    login: async (username, password) => {
        try {
            const { accessToken, refreshToken, user } = await api.login(username, password);
            
            // Store token and user in local storage
            localStorage.setItem('token', accessToken);
            localStorage.setItem('refreshToken', refreshToken);
            localStorage.setItem('user', JSON.stringify(user));
            
            return { accessToken, refreshToken, user };
        } catch (error) {
            console.error('Login error:', error);
            throw error;
        }
    },
    
    // Register new user
    signup: async (userData) => {
        try {
            let response = await api.signup(userData)

            // Store token and user in local storage
            localStorage.setItem('refreshToken', response.refreshToken);
            localStorage.setItem('token', response.accessToken);
            localStorage.setItem('user', JSON.stringify(response.user));

            return response;
        } catch (error) {
            console.error('Signup error:', error);
            throw error;
        }
    },
    
    // Logout user
    logout: () => {            
        // Remove token and user from local storage
        localStorage.removeItem('token');
        localStorage.removeItem('refreshToken');
        localStorage.removeItem('user');
        
        return { isAuthenticated: false, token: null, user: null };
    },
    
    // Check if user is authenticated
    isAuthenticated: async () => {
        
        let u = localStorage.getItem('user');
        const usr = await api.getUser(JSON.parse(u).id);
        localStorage.setItem('usr', JSON.stringify(usr));
        
        const token = localStorage.getItem('token');
        return !!token;
    },
    
    // Get current user
    getUser: () => {
        const user = localStorage.getItem('user');
        return user ? JSON.parse(user) : null;
    },
    
    // Get auth token
    getToken: () => {
        return localStorage.getItem('token');
    }
};

// Update auth UI elements
const updateAuthUI = () => {
    const authContainer = document.getElementById('authContainer');
    if (!authContainer) return;
    
    const isAuthenticated = auth.isAuthenticated();
    const user = auth.getUser();
    
    // Clear auth container
    authContainer.innerHTML = '';
    
    if (isAuthenticated && user) {
        // User is logged in, show user info and logout button
        const userInfo = document.createElement('div');
        userInfo.className = 'user-info';
        userInfo.innerHTML = `
            <span>${user.username}</span>
            <button id="logoutBtn" class="btn">Logout</button>
        `;
        
        // Add new post button if on main page
        // if (window.location.pathname === '/' || window.location.pathname.endsWith('index.html')) {
        //     const newPostBtn = document.createElement('button');
        //     newPostBtn.id = 'newPostBtn';
        //     newPostBtn.className = 'btn new-post-btn';
        //     newPostBtn.textContent = 'New Post';
        //     userInfo.appendChild(newPostBtn);
            
        //     // Add event listener for new post button
        //     setTimeout(() => {
        //         const newPostBtn = document.getElementById('newPostBtn');
        //         if (newPostBtn) {
        //             newPostBtn.addEventListener('click', () => {
        //                 const newPostModal = document.getElementById('newPostModal');
        //                 newPostModal.style.display = 'block';
        //             });
        //         }
        //     }, 0);
        // }
        
        authContainer.appendChild(userInfo);
        
        // Add event listener for logout button
        setTimeout(() => {
            const logoutBtn = document.getElementById('logoutBtn');
            if (logoutBtn) {
                logoutBtn.addEventListener('click', () => {
                    auth.logout();
                    updateAuthUI();
                    // Reload current page after logout
                    window.location.reload();
                });
            }
        }, 0);
    } else {
        // User is not logged in, show login and signup buttons
        const loginBtn = document.createElement('button');
        loginBtn.id = 'loginBtn';
        loginBtn.className = 'btn';
        loginBtn.textContent = 'Login';
        
        const signupBtn = document.createElement('button');
        signupBtn.id = 'signupBtn';
        signupBtn.className = 'btn';
        signupBtn.textContent = 'Sign Up';
        
        authContainer.appendChild(loginBtn);
        authContainer.appendChild(signupBtn);
        
        // Add event listeners for login and signup buttons
        setTimeout(() => {
            const loginBtn = document.getElementById('loginBtn');
            const signupBtn = document.getElementById('signupBtn');
            const loginModal = document.getElementById('loginModal');
            const signupModal = document.getElementById('signupModal');
            
            if (loginBtn && loginModal) {
                loginBtn.addEventListener('click', () => {
                    loginModal.style.display = 'block';
                });
            }
            
            if (signupBtn && signupModal) {
                signupBtn.addEventListener('click', () => {
                    signupModal.style.display = 'block';
                });
            }
        }, 0);
    }
};

// Initialize modals
const initModals = () => {
    // Get all close buttons and modals
    const closeButtons = document.querySelectorAll('.close');
    const modals = document.querySelectorAll('.modal');
    const showLoginBtn = document.getElementById('showLoginBtn');
    const showSignupBtn = document.getElementById('showSignupBtn');
    const loginModal = document.getElementById('loginModal');
    const signupModal = document.getElementById('signupModal');
    const loginForm = document.getElementById('loginForm');
    const signupForm = document.getElementById('signupForm');
    const newPostForm = document.getElementById('newPostForm');
    
    // Add event listeners for close buttons
    closeButtons.forEach(button => {
        button.addEventListener('click', () => {
            modals.forEach(modal => {
                modal.style.display = 'none';
            });
        });
    });
    
    // Close modal when clicking outside of it
    window.addEventListener('click', (event) => {
        modals.forEach(modal => {
            if (event.target === modal) {
                modal.style.display = 'none';
            }
        });
    });
    
    // Switch between login and signup modals
    if (showLoginBtn && loginModal && signupModal) {
        showLoginBtn.addEventListener('click', (event) => {
            event.preventDefault();
            signupModal.style.display = 'none';
            loginModal.style.display = 'block';
        });
    }
    
    if (showSignupBtn && loginModal && signupModal) {
        showSignupBtn.addEventListener('click', (event) => {
            event.preventDefault();
            loginModal.style.display = 'none';
            signupModal.style.display = 'block';
        });
    }
    
    // Login form submission
    if (loginForm) {
        loginForm.addEventListener('submit', async (event) => {
            event.preventDefault();
            
            const username = document.getElementById('loginUsername').value;
            const password = document.getElementById('loginPassword').value;
            
            try {
                await auth.login(username, password);
                // Close modal and update UI
                loginModal.style.display = 'none';
                updateAuthUI();
                // Reload current page after login
                window.location.reload();
            } catch (error) {
                alert('Login failed. Please check your credentials and try again.');
            }
        });
    }
    
    // Signup form submission
    if (signupForm) {
        signupForm.addEventListener('submit', async (event) => {
            event.preventDefault();
            
            const userData = {
                username: document.getElementById('signupUsername').value,
                email: document.getElementById('signupEmail').value,
                password: document.getElementById('signupPassword').value,
                avatarUrl: document.getElementById('avatarImage').value
            };
            
            try {
                await auth.signup(userData);

                // Close modal and update UI
                signupModal.style.display = 'none';
                updateAuthUI();
                // Reload current page after signup
                window.location.reload();
            } catch (error) {
                alert('Sign up failed. Please try again with different credentials.');
            }
        });

        // const form = document.getElementById('signupForm');

        // form.addEventListener('submit', async function(e) {
        //     e.preventDefault();
            
        //     const formData = new FormData(this);
            
        //     try {
        //         const response = await fetch('/api/users/signup', {
        //             method: 'POST',
        //             body: formData // FormData handles the multipart/form-data encoding
        //         });
                
        //         const result = await response.json();
        //         // Handle response
                
        //     } catch (error) {
        //         console.error('Error:', error);
        //     }
        // });
    }
    
    // New post form submission
    if (newPostForm) {
        newPostForm.addEventListener('submit', async (event) => {
            event.preventDefault();
            
            const postData = {
                title: document.getElementById('postTitle').value,
                content: document.getElementById('postContent').value
            };
            
            try {
                const token = auth.getToken();
                await api.createPost(postData, token);
                // Close modal and reload page to show new post
                document.getElementById('newPostModal').style.display = 'none';
                window.location.reload();
            } catch (error) {
                alert('Failed to create post. Please try again.');
            }
        });
    }
};

// Initialize theme toggle
const initThemeToggle = () => {
    const themeToggle = document.getElementById('themeToggle');
    if (!themeToggle) return;
    
    // Check for saved theme preference or default to light
    const currentTheme = localStorage.getItem('theme') || 'light';
    document.body.className = currentTheme === 'dark' ? 'dark-mode' : 'light-mode';
    
    // Toggle theme when button is clicked
    themeToggle.addEventListener('click', () => {
        if (document.body.className === 'light-mode') {
            document.body.className = 'dark-mode';
            localStorage.setItem('theme', 'dark');
        } else {
            document.body.className = 'light-mode';
            localStorage.setItem('theme', 'light');
        }
    });
};

// Initialize auth on page load
document.addEventListener('DOMContentLoaded', () => {
    initThemeToggle();
    updateAuthUI();
    initModals();
});

// Add this to your auth.js or app.js file
document.addEventListener('DOMContentLoaded', function() {
    const avatarInput = document.getElementById('avatarImage');
    const fileNameDisplay = document.getElementById('fileNameDisplay');
    const avatarPreview = document.getElementById('avatarPreview');
    
    if (avatarInput) {
        avatarInput.addEventListener('change', function() {
            if (this.files && this.files[0]) {
                // Display filename
                fileNameDisplay.textContent = this.files[0].name;
                
                // Show preview
                const reader = new FileReader();
                reader.onload = function(e) {
                    avatarPreview.style.backgroundImage = `url(${e.target.result})`;
                    avatarPreview.style.display = 'block';
                }
                reader.readAsDataURL(this.files[0]);
            } else {
                fileNameDisplay.textContent = 'No file chosen';
                avatarPreview.style.display = 'none';
                avatarPreview.style.backgroundImage = 'none';
            }
        });
    }
});