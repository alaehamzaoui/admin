"use client"
import React, { useEffect } from 'react'
import * as z from "zod"
import { Button } from "@/components/ui/button"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from 'react-hook-form'
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import Link from 'next/link'
import {BACKEND_URL} from '@/loadEnv';
import router from 'next/router'


const signInSchema = z.object({
    email: z.string().email("Email must be valid."),
    password: z.string().min(6, "Password Should have atleast 6 characters."),
})

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
    }, [router]);*/
    const form = useForm<z.infer<typeof signInSchema>>({
        resolver: zodResolver(signInSchema),
        defaultValues: {
          email: "",
          password: "",
        },
    })
    async function onSubmit(data: z.infer<typeof signInSchema>) {
        try {
                const loginResponse = await fetch(BACKEND_URL + "/auth/login", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        email: data.email,
                        password: data.password,
                    }),
                });
                if (loginResponse.ok) {
                    const loginData = await loginResponse.json();
                    alert("login token : " + loginData.token);
                    localStorage.setItem("authToken", loginData.token);
                    window.location.href= "/home";

                } else {
                    console.error("Failed to login", loginResponse.status);
                }
        } catch (error) {
            console.error('Failed to submit data', error);
        }
    }
  return (
    <>
        <div className="signUpWrapper">
            <div className="formWrapper">
                <div className="left">
                    <h3 className="title">Hallo, Freunde!</h3>
                    <p>Geben Sie Ihre persönlichen Daten ein und starten Sie Ihre Reise mit uns</p>
                    <Link href={"/signup"}>
                    <Button className='border-zinc-500 text-white-300 hover:border-zinc-200 hover:text-white-100 hover:bg-teal-600 transition-colors border rounded-full px-8 bg-teal-600'>Sign Up</Button>
                    </Link>
                </div>
                <div className="right">
                    <h3 className='text-center text-2xl font-semibold'>Hier anmelden</h3>
                    <Form {...form}>
                        <form onSubmit={form.handleSubmit(onSubmit)}>
                            <FormField
                                control={form.control}
                                name="email"
                                render={({ field }) => (
                                    <FormItem className='space-y-0 mb-2'>
                                        <FormLabel>Email</FormLabel>
                                        <FormControl>
                                            <Input placeholder="admin@example.com" {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="password"
                                render={({ field }) => (
                                    <FormItem className='space-y-0 mb-2'>
                                        <FormLabel>Password</FormLabel>
                                        <FormControl>
                                            <Input placeholder="********" type='password' {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <Button type="submit" className='w-full bg-teal-600 text-zinc-200  hover:border-zinc-200 hover:text-zinc-100  hover:bg-teal-600' >Submit</Button>
                        </form>
                    </Form>
                </div>
            </div>
        </div>
    </>
  )
}

export default Page