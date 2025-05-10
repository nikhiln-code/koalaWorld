"use client";

import { useEffect, useRef } from "react";

const KOALA_IMAGES = [
  "/koala/k1.png",
  "/koala/k2.png",
  "/koala/k3.png",
  "/koala/k4.png",
  "/koala/k5.png",
  "/koala/k6.png",
];

export default function BackgroundEffect() {
  const canvasRef = useRef<HTMLCanvasElement | null>(null);
  const mouseRef = useRef({ x: 0, y: 0 });
  // Koala image particles
  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const ctx = canvas.getContext("2d");
    if (!ctx) return;

    const width = (canvas.width = window.innerWidth);
    const height = (canvas.height = window.innerHeight);

    const loadedImages: HTMLImageElement[] = [];
    let loadedCount = 0;

    KOALA_IMAGES.forEach((src, index) => {
      const img = new Image();
      img.src = src;
      img.onload = () => {
        loadedImages[index] = img;
        loadedCount++;
        if (loadedCount === KOALA_IMAGES.length) startAnimation();
      };
    });

    const particles = Array.from({ length: 100 }, () => ({
      x: Math.random() * width,
      y: Math.random() * height,
      dx: (Math.random() - 0.5) * 1.2,
      dy: (Math.random() - 0.5) * 1.2,
      size: 32 + Math.random() * 16,
      image: KOALA_IMAGES[Math.floor(Math.random() * KOALA_IMAGES.length)],
    }));

    function startAnimation() {
      function animate() {
        ctx.clearRect(0, 0, width, height);

        particles.forEach((p) => {
          // Repel from mouse
          const dx = p.x - mouseRef.current.x;
          const dy = p.y - mouseRef.current.y;
          const dist = Math.sqrt(dx * dx + dy * dy);
          if (dist < 100) {
            p.x += dx / dist;
            p.y += dy / dist;
          } else {
            p.x += p.dx;
            p.y += p.dy;
          }

          // Wrap around edges
          if (p.x < -p.size) p.x = width;
          if (p.x > width) p.x = -p.size;
          if (p.y < -p.size) p.y = height;
          if (p.y > height) p.y = -p.size;

          const img = loadedImages[KOALA_IMAGES.indexOf(p.image)];
          if (img) ctx.drawImage(img, p.x, p.y, p.size, p.size);
        });

        requestAnimationFrame(animate);
      }
      animate();
    }

    const handleMouseMove = (e: MouseEvent) => {
      mouseRef.current.x = e.clientX;
      mouseRef.current.y = e.clientY;
    };

    window.addEventListener("mousemove", handleMouseMove);
    return () => window.removeEventListener("mousemove", handleMouseMove);
  }, []);

  return (
    <>
      <img
        src="/jungle-bg.jpg"
        alt="jungle background"
        className="absolute inset-0 w-full h-full object-cover opacity-40 blur-xs z-0"
      />
      <canvas
        ref={canvasRef}
        className="absolute inset-0 z-0 opacity-40 min-h-[170vh] max-w-full"
        style={{ pointerEvents: "none" }}
      />
    </>
  );
}
