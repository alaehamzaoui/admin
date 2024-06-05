"use client"
import React, { useEffect } from 'react'
import { Button } from "@/components/ui/button"
import Link from 'next/link'
import { BACKEND_URL } from '@/loadEnv';
import router from 'next/router'

const Page = () => {
    /*
    useEffect(() => {
        const checkUserLoggedIn = async () => {
            const authToken = localStorage.getItem('authToken'); // Assuming the API token is stored with this key
                if (!authToken) {
                    console.error('Failed to get token from localstorage :');
                    router.push('/signin');
                } else {
                    try {
                        const headers = new Headers();
                        headers.append('X-API-Key', authToken ?? ''); // Use the API token for the request
            
                        const response = await fetch(`${BACKEND_URL}/system/user`, {
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
                    //todo : check if token is valid http://localhost:4664/system/user mit X-API-Key = token
                    router.push('/home');
                    }
                    catch (error) {
                        console.error('Failed to submit data', error);
                    }
                }
        };
    checkUserLoggedIn();
    }, [router]);
    */
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
                alert(url);
                window.location.href = "http://"+url; // Redirect to the URL with the token
            }
            //todo : check if token is valid http://localhost:4664/system/user mit X-API-Key = token
            router.push('/home');
            }
          catch (error) {
            console.error('Failed to submit data', error);
          }
    }

    const handleUpdatePassword = () => {
        
    }

    const handleLogout = () => {
        localStorage.removeItem('authToken');
        window.location.href = '/signin';
    }

    return (
        <>
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

            `}</style>
        </>
    )
}

export default Page
