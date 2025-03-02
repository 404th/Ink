// Main App JavaScript - For index.html

// Load and display posts
const loadPosts = async () => {
    const postsContainer = document.getElementById('postsContainer');
    if (!postsContainer) return;
    
    try {
        // Clear loading indicator
        postsContainer.innerHTML = '';
        
        // Fetch posts from API
        const posts = await api.getPosts();
        
        if (posts.length === 0) {
            postsContainer.innerHTML = '<p class="no-posts">No posts found.</p>';
            return;
        }
        
        // Sort posts by creation date (newest first)
        posts.sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt));
        
        // Create post cards
        posts.forEach(post => {
            const postCard = document.createElement('div');
            postCard.className = 'post-card';
            postCard.setAttribute('data-id', post.id);
            
            postCard.innerHTML = `
                <h2>${post.title}</h2>
                <div class="post-meta">
                    <span>by ${post.author}</span>
                    <span>Rating: ${post.rating}</span>
                </div>
            `;
            
            // Add click event to navigate to post
            postCard.addEventListener('click', () => {
                window.location.href = `post.html?id=${post.id}`;
            });
            
            postsContainer.appendChild(postCard);
        });
    } catch (error) {
        console.error('Error loading posts:', error);
        postsContainer.innerHTML = '<p class="error">Failed to load posts. Please try again later.</p>';
    }
};

// Add new post button if user is logged in
const addNewPostButton = () => {
    if (auth.isAuthenticated()) {
        const postsContainer = document.getElementById('postsContainer');
        if (!postsContainer) return;
        
        const newPostBtn = document.createElement('button');
        newPostBtn.className = 'btn new-post-btn';
        newPostBtn.textContent = 'Create New Post';
        newPostBtn.addEventListener('click', () => {
            const newPostModal = document.getElementById('newPostModal');
            if (newPostModal) {
                newPostModal.style.display = 'block';
            }
        });
        
        // Insert before the postsContainer
        postsContainer.parentNode.insertBefore(newPostBtn, postsContainer);
    }
};

// Initialize index page
document.addEventListener('DOMContentLoaded', () => {
    loadPosts();
    addNewPostButton();
});