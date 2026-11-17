/**
 * Liam's Planet Simulator - Clean & Simple
 * - Polls GET /state every 50ms for body positions
 * - Renders celestial bodies & trails onto 2D Canvas
 * - Updates inspector sidebar for selected planet
 */

const canvas = document.getElementById('spaceCanvas');
const ctx = canvas.getContext('2d');

const statusEl = document.getElementById('status');
const nameEl = document.getElementById('planet-name');
const dotEl = document.getElementById('inspector-dot');
const posEl = document.getElementById('stat-pos');
const velEl = document.getElementById('stat-vel');
const speedEl = document.getElementById('stat-speed');
const radiusEl = document.getElementById('stat-radius');
const planetListEl = document.getElementById('planet-list');

// Simulation State
let bodies = [];
let selectedIndex = 0;
const trails = new Map(); // name -> array of {x, y}
const MAX_TRAIL = 60;

// High-DPI Canvas Resizing
function resize() {
  const dpr = window.devicePixelRatio || 1;
  canvas.width = window.innerWidth * dpr;
  canvas.height = window.innerHeight * dpr;
  ctx.resetTransform?.();
  ctx.scale(dpr, dpr);
}
window.addEventListener('resize', resize);
resize();

// 1. Polling GET /state
async function fetchState() {
  try {
    const res = await fetch('/state');
    if (res.ok) {
      const data = await res.json();
      // Accepts either an array [ {...} ] or an object { bodies: [ {...} ] }
      bodies = Array.isArray(data) ? data : (data.bodies || []);
      statusEl.textContent = `Live: ${bodies.length} bodies tracked`;
      statusEl.style.color = '#34d399';
    } else {
      statusEl.textContent = `GET /state returned HTTP ${res.status}`;
      statusEl.style.color = '#0bf580ff';
    }
  } catch (err) {
    statusEl.textContent = 'Waiting for backend at GET /state...';
    statusEl.style.color = '#9ca3af';
  }
}
setInterval(fetchState, 50);
fetchState();

// 2. Click to select a planet
canvas.addEventListener('click', (e) => {
  const cx = window.innerWidth / 2;
  const cy = window.innerHeight / 2;
  const clickX = e.clientX - cx;
  const clickY = e.clientY - cy;

  bodies.forEach((b, idx) => {
    const dist = Math.hypot(b.x - clickX, b.y - clickY);
    if (dist <= Math.max(12, b.radius + 4)) {
      selectedIndex = idx;
    }
  });
});

// 3. Render Loop
function render() {
  const w = window.innerWidth;
  const h = window.innerHeight;
  const cx = w / 2;
  const cy = h / 2;

  // Clear space background
  ctx.fillStyle = '#0b0f19';
  ctx.fillRect(0, 0, w, h);

  // Center crosshair
  ctx.strokeStyle = 'rgba(255, 255, 255, 0.08)';
  ctx.beginPath();
  ctx.moveTo(cx - 10, cy);
  ctx.lineTo(cx + 10, cy);
  ctx.moveTo(cx, cy - 10);
  ctx.lineTo(cx, cy + 10);
  ctx.stroke();

  // Draw Trails and Bodies
  bodies.forEach((b, idx) => {
    const color = b.color || '#38bdf8';
    const screenX = cx + b.x;
    const screenY = cy + b.y;

    // Track trail
    if (!trails.has(b.name)) trails.set(b.name, []);
    const trail = trails.get(b.name);
    trail.push({ x: screenX, y: screenY });
    if (trail.length > MAX_TRAIL) trail.shift();

    // Render trail line
    if (trail.length > 1) {
      ctx.strokeStyle = color;
      ctx.globalAlpha = 0.35;
      ctx.lineWidth = 1.5;
      ctx.beginPath();
      ctx.moveTo(trail[0].x, trail[0].y);
      for (let i = 1; i < trail.length; i++) {
        ctx.lineTo(trail[i].x, trail[i].y);
      }
      ctx.stroke();
      ctx.globalAlpha = 1.0;
    }

    // Render planet body
    ctx.fillStyle = color;
    ctx.beginPath();
    ctx.arc(screenX, screenY, b.radius || 5, 0, Math.PI * 2);
    ctx.fill();

    // Planet label
    if (b.name) {
      ctx.fillStyle = '#9ca3af';
      ctx.font = '11px sans-serif';
      ctx.textAlign = 'center';
      ctx.fillText(b.name, screenX, screenY + (b.radius || 5) + 12);
    }

    // Selection ring
    if (idx === selectedIndex) {
      ctx.strokeStyle = '#ffffff';
      ctx.lineWidth = 1.5;
      ctx.beginPath();
      ctx.arc(screenX, screenY, (b.radius || 5) + 4, 0, Math.PI * 2);
      ctx.stroke();
    }
  });

  // Update Inspector Panel
  updateInspector();

  requestAnimationFrame(render);
}

// 4. Update Inspector Sidebar
function updateInspector() {
  if (bodies.length === 0) return;

  if (selectedIndex >= bodies.length) selectedIndex = 0;
  const b = bodies[selectedIndex];

  nameEl.textContent = b.name || `Body #${selectedIndex + 1}`;
  dotEl.style.backgroundColor = b.color || '#38bdf8';
  posEl.textContent = `(${Math.round(b.x)}, ${Math.round(b.y)})`;

  const vx = b.vx || 0;
  const vy = b.vy || 0;
  velEl.textContent = `(${vx.toFixed(1)}, ${vy.toFixed(1)})`;
  speedEl.textContent = `${Math.hypot(vx, vy).toFixed(1)}`;
  radiusEl.textContent = `${b.radius || '-'}`;

  // Update list buttons if count changed
  if (planetListEl.children.length !== bodies.length) {
    planetListEl.innerHTML = '';
    bodies.forEach((body, idx) => {
      const btn = document.createElement('button');
      btn.className = `planet-btn ${idx === selectedIndex ? 'selected' : ''}`;
      btn.innerHTML = `<span style="width:8px; height:8px; border-radius:50%; background:${body.color || '#38bdf8'}; display:inline-block;"></span> ${body.name || `Body ${idx+1}`}`;
      btn.addEventListener('click', () => {
        selectedIndex = idx;
      });
      planetListEl.appendChild(btn);
    });
  } else {
    Array.from(planetListEl.children).forEach((el, idx) => {
      el.classList.toggle('selected', idx === selectedIndex);
    });
  }
}

requestAnimationFrame(render);
