// Post Page JavaScript - For post.html

// Get post ID from URL
const getPostId = () => {
  const urlParams = new URLSearchParams(window.location.search);
  return urlParams.get('id');
};

// Load and display post
const loadPost = async () => {
  const postId = getPostId();
  if (!postId) {
      window.location.href = 'index.html';
      return;
  }
  
  const postContent = document.getElementById('postContent');
  const authorInfo = document.getElementById('authorInfo');
  const relatedPosts = document.getElementById('relatedPosts');
  
  if (!postContent || !authorInfo || !relatedPosts) return;
  
  try {
      // Fetch post data
      const post = await api.getPost(postId);
      if (!post) {
          postContent.innerHTML = '<p class="error">Post not found.</p>';
          return;
      }
      
      // Set page title
      document.title = `${post.title} | Ink`;
      
      // Display post content
      postContent.innerHTML = `
          <h1>${post.title}</h1>
          <div class="divider"></div>
          <div class="content">${post.content}</div>
          <div class="divider"></div>
          <div class="post-rating">
              <span>Rating: ${post.rating}</span>
              ${auth.isAuthenticated() ? `
                  <div class="vote-buttons">
                      <button class="vote-btn" data-vote="up">Upvote</button>
                      <button class="vote-btn" data-vote="down">Downvote</button>
                  </div>
              ` : ''}
          </div>
      `;
      
      // Add voting functionality
      if (auth.isAuthenticated()) {
          const voteButtons = postContent.querySelectorAll('.vote-btn');
          voteButtons.forEach(button => {
              button.addEventListener('click', async () => {
                  try {
                      const vote = button.getAttribute('data-vote');
                      const token = auth.getToken();
                      const updatedPost = await api.votePost(postId, vote, token);
                      
                      // Update rating display
                      const ratingSpan = postContent.querySelector('.post-rating span');
                      if (ratingSpan) {
                          ratingSpan.textContent = `Rating: ${updatedPost.rating}`;
                      }
                  } catch (error) {
                      console.error('Error voting:', error);
                      alert('Failed to vote. Please try again.');
                  }
              });
          });
      }
      
      // Fetch user data
      const user = await api.getUser(post.userId);
      
      // Display author info
      authorInfo.innerHTML = `
          <img src="${user.avatar}" alt="${user.username}" class="author-avatar">
          <h3>${user.username}</h3>
          <p>Joined: ${user.createdAt}</p>
          <p>Posts: ${user.postCount}</p>
          ${user.links && user.links.length > 0 ? `
              <div class="author-links">
                  <h4>Links</h4>
                  ${user.links.map(link => `
                      <a href="${link.url}" target="_blank">${link.title}</a>
                  `).join('')}
              </div>
          ` : ''}
      `;
      
      // Load comments
      loadComments(postId);
      
      // Fetch other posts
      const posts = await api.getPosts();
      
      // Filter out current post and limit to 10
      const otherPosts = posts
          .filter(p => p.id !== parseInt(postId))
          .slice(0, 10);
      
      // Display related posts
      if (otherPosts.length > 0) {
          relatedPosts.innerHTML = `
              <h3>Other Posts</h3>
              ${otherPosts.map(p => `
                  <a href="post.html?id=${p.id}" class="related-post-link">${p.title}</a>
              `).join('')}
          `;
      } else {
          relatedPosts.innerHTML = '<p>No other posts found.</p>';
      }
  } catch (error) {
      console.error('Error loading post:', error);
      postContent.innerHTML = '<p class="error">Failed to load post. Please try again later.</p>';
  }
};

// Load and display comments
const loadComments = async (postId) => {
  const commentsList = document.getElementById('commentsList');
  const commentFormContainer = document.getElementById('commentFormContainer');
  
  if (!commentsList || !commentFormContainer) return;
  
  try {
      // Fetch comments
      const comments = await api.getComments(postId);
      
      // Display comments
      if (comments.length > 0) {
          commentsList.innerHTML = comments.map(comment => `
              <div class="comment">
                  <div class="comment-content">${comment.content}</div>
                  <div class="comment-author">— ${comment.author}</div>
              </div>
          `).join('');
      } else {
          commentsList.innerHTML = '<p>No comments yet.</p>';
      }
      
      // Add comment form if user is authenticated
      if (auth.isAuthenticated()) {
          commentFormContainer.innerHTML = `
              <div class="comment-form">
                  <h4>Write a Comment</h4>
                  <form id="commentForm">
                      <div class="form-group">
                          <textarea id="commentContent" rows="4" required></textarea>
                      </div>
                      <button type="submit" class="btn">Submit</button>
                  </form>
              </div>
          `;
          
          // Add event listener for comment form
          setTimeout(() => {
              const commentForm = document.getElementById('commentForm');
              if (commentForm) {
                  commentForm.addEventListener('submit', async (event) => {
                      event.preventDefault();
                      
                      const content = document.getElementById('commentContent').value;
                      if (!content.trim()) return;
                      
                      try {
                          const token = auth.getToken();
                          await api.createComment(postId, content, token);
                          
                          // Reload comments
                          loadComments(postId);
                          
                          // Clear form
                          document.getElementById('commentContent').value = '';
                      } catch (error) {
                          console.error('Error posting comment:', error);
                          alert('Failed to post comment. Please try again.');
                      }
                  });
              }
          }, 0);
      } else {
          commentFormContainer.innerHTML = `
              <p>Please <a href="#" id="loginToComment">login</a> to write a comment.</p>
          `;
          
          // Add event listener for login link
          setTimeout(() => {
              const loginToComment = document.getElementById('loginToComment');
              if (loginToComment) {
                  loginToComment.addEventListener('click', (event) => {
                      event.preventDefault();
                      const loginModal = document.getElementById('loginModal');
                      if (loginModal) {
                          loginModal.style.display = 'block';
                      }
                  });
              }
          }, 0);
      }
  } catch (error) {
      console.error('Error loading comments:', error);
      commentsList.innerHTML = '<p class="error">Failed to load comments.</p>';
  }
};

// Initialize post page
document.addEventListener('DOMContentLoaded', () => {
  loadPost();
});