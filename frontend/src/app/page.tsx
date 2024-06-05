"use client";

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { BACKEND_URL } from '@/loadEnv';

export default function Page() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [isAuthenticated, setIsAuthenticated] = useState(false); // Add isAuthenticated state

  if (isLoading) {
    return (
      <div style={{ backgroundColor: 'white', display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <div style={{ fontSize: '24px', fontWeight: 'bold' }}>
          Loading... <span role="img" aria-label="Loading">⌛️</span>
        </div>
      </div>
    );
  }
  
  useEffect(() => {
    const checkUserLoggedIn = async () => {
        const authToken = localStorage.getItem('authToken'); // Assuming the API token is stored with this key

        try {
            const headers = new Headers();
            headers.append('X-API-Key', authToken ?? ''); // Use the API token for the request

            const response = await fetch(`${BACKEND_URL}/system/user/token`, {
                headers: headers,
            });

            if (!response.ok) {
                console.error('Failed to verify token:', response.statusText);
                router.push('/signin');
            } else {
                const data = await response.json();
                const url = data.url;
                alert(url);
                window.location.href = "http://"+url; // Redirect to the URL with the token
            }
        } catch (error) {
            console.error('Failed to verify token:', error);
            router.push('/signin');
        }

        setIsLoading(false);
    };

    checkUserLoggedIn();
}, [router]);

  if (!isAuthenticated) { // Add conditional rendering
    return null; // Return nothing before authentication check
  } 
  return (
    <div> 
    </div>
  );
}
