"use client"
import React, { useState, useEffect } from 'react'
import { Button } from "@/components/ui/button"
import { BACKEND_URL } from '@/loadEnv';

const Page = () => {
    const [loading, setLoading] = useState(true);
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    useEffect(() => {
        const checkUserLoggedIn = async () => {
            const authToken = localStorage.getItem('authToken') || ""; // Assuming the API token is stored with this key
            if (!authToken) {
                console.error('Failed to get token from localstorage :');
                window.location.href = "/signin";
            } else {
                try {
                    const headers = new Headers();
                    headers.append('X-API-Key', authToken); // Use the API token for the request
                    const response = await fetch(`${BACKEND_URL}/system/user`, {
                        headers: headers,
                    });
                    if (!response.ok) {
                        console.error('Failed to verify token:', response.statusText);
                        window.location.href = "/signin";
                    } else {
                        setIsAuthenticated(true);
                    }
                } catch (error) {
                    console.error('Failed to submit data', error);
                    window.location.href = "/signin";
                }
            }
            setLoading(false);
        };
        checkUserLoggedIn();
    }, []);

    const handleWorkAdventure = async () => {
        try {
            const authToken = localStorage.getItem('authToken'); // Assuming the API token is stored with this key
            const headers = new Headers();
            headers.append('X-API-Key', authToken ?? ''); // Use the API token for the request

            const response = await fetch(`${BACKEND_URL}/system/user/token`, {
                headers: headers,
            });
            if (!response.ok) {
                console.error('Failed to verify token:', response.statusText);
            } else {
                const data = await response.json();
                const url = data.url;
                window.location.href = "http://" + url; // Redirect to the URL with the token
            }
        }
        catch (error) {
            console.error('Failed to submit data', error);
        }
    }

    const handleUpdatePassword = () => {
        // Handle password update
    }

    const handleLogout = () => {
        localStorage.removeItem('authToken');
        window.location.href = '/signin';
    }

    if (loading) {
        return (
            <div className="loading-wrapper">
                <h3>Loading...</h3>
            </div>
        );
    }

    return (
        <div className={isAuthenticated ? "visible" : "hidden"}>
            <div className="bg-wrapper">
                <div className="signUpWrapper">
                    <div className="formWrapper">
                        <div className="left">
                            <h3 className='text-center text-4xl font-semibold'>Willkommen</h3>
                            <h3 className='text-center text-4xl font-semibold'> username</h3>
                        </div>
                        <div className="right">
                            <div className="space-y-8 mt-8">
                                <Button className='w-full bg-teal-600 text-zinc-200 hover:border-zinc-200 hover:text-zinc-100 hover:bg-teal-600' onClick={handleWorkAdventure}>Go to WorkAdventure</Button>
                                <Button className='w-full bg-teal-600 text-zinc-200 hover:border-zinc-200 hover:text-zinc-100 hover:bg-teal-600' onClick={handleUpdatePassword}>Change Password</Button>
                                <Button className='w-full bg-teal-600 text-zinc-200 hover:border-zinc-200 hover:text-zinc-100 hover:bg-teal-600' onClick={handleLogout}>Log Out</Button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <style jsx>{`
                .loading-wrapper {
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    font-size: 24px;
                    color: #333;
                }
                .hidden {
                    display: none;
                }
                .visible {
                    display: block;
                }
            `}</style>
        </div>
    )
}

export default Page
