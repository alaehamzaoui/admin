"use client";

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { BACKEND_URL } from '@/loadEnv';

export default function Page() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(true); // Start with true to show the loading state initially

  useEffect(() => {
    const checkUserLoggedIn = async () => {
      const authToken = localStorage.getItem('authToken'); // Assuming the API token is stored with this key

      try {
        const headers = new Headers();
        headers.append('X-API-Key', authToken ?? ''); // Use the API token for the request

        const response = await fetch(`${BACKEND_URL}/system/user/`, {
          headers: headers,
        });
        if (!response.ok) {
          console.error('Failed to verify token:', response.statusText);
          window.location.href = '/signin';
          return; // Add return to stop further execution
        }
        const data = await response.json();
        if (data.email) {
          window.location.href = '/home';
        } else {
          window.location.href = '/signin';
        }
      } catch (error) {
        console.error('Failed to verify token:', error);
        window.location.href = '/signin';
      }
      setIsLoading(false);
    };
    checkUserLoggedIn();
  }, [router]);

  if (isLoading) {
    return (
      <div style={{ backgroundColor: 'white', display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <div style={{ fontSize: '24px', fontWeight: 'bold' }}>
          Loading... <span role="img" aria-label="Loading">⌛️</span>
        </div>
      </div>
    );
  }

  return (
    <div>
      {/* Your main content here */}
    </div>
  );
}
