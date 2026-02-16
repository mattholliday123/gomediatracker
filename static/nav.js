// Shared Navigation Component
// This script can be used to dynamically inject navigation into pages

function createNavigation(activePage = '') {
  const nav = document.createElement('nav');
  nav.className = 'nav-bar';
  
  nav.innerHTML = `
    <div class="container">
      <a href="home.html" class="nav-brand">
        <strong>Media Collection</strong>
      </a>
      <ul class="nav-links">
        <li><a href="all.html" ${activePage === 'all' ? 'class="active"' : ''}>All Media</a></li>
        <li><a href="books.html" ${activePage === 'books' ? 'class="active"' : ''}>Books</a></li>
        <li><a href="movies.html" ${activePage === 'movies' ? 'class="active"' : ''}>Movies</a></li>
        <li><a href="music.html" ${activePage === 'music' ? 'class="active"' : ''}>Music</a></li>
        <li><a href="video_games.html" ${activePage === 'video_games' ? 'class="active"' : ''}>Video Games</a></li>
      </ul>
    </div>
  `;
  
  return nav;
}

// Auto-detect active page based on current URL
function getActivePage() {
  const path = window.location.pathname;
  const filename = path.split('/').pop() || 'home.html';
  
  if (filename === 'home.html' || filename === 'index.html' || filename === '') {
    return 'home';
  }
  
  return filename.replace('.html', '');
}

// Initialize navigation if nav-bar element exists
document.addEventListener('DOMContentLoaded', () => {
  const existingNav = document.querySelector('.nav-bar');
  if (!existingNav) {
    const body = document.body;
    const firstChild = body.firstElementChild;
    const nav = createNavigation(getActivePage());
    
    if (firstChild) {
      body.insertBefore(nav, firstChild);
    } else {
      body.appendChild(nav);
    }
  }
});
