import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import React from "react";
import { useNavigate } from "react-router";

export function LoginForm({
  className,
  username, password, setUsername, setPassword, user, setUser,
  ...props
}: React.ComponentProps<"div">) {

  const navigate = useNavigate();

  const login = async() => {
    const response = await fetch("http://localhost:8000/login", {
      method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({username: username, password: password}),
    })
    const res = response.json();
    console.log(res);
    setUser(localStorage.setItem("user", username));
    navigate("/dashboard");
  }

  return (
    <div className={cn("flex flex-col gap-6 w-4/12", className)} {...props}>
      <Card>
        <CardHeader className="text-center">
          <CardTitle className="text-xl">Welcome</CardTitle>
        </CardHeader>
        <CardContent>
            <div className="grid gap-6">
              <div className="grid gap-6">
                <div className="grid gap-3">
                  <Label htmlFor="email">Username</Label>
                  <Input
                    id="email"
                    placeholder="mahesh dhalle"
                    required
                    onChange={(e) => setUsername(e.target.value)}
                  />
                </div>
                <div className="grid gap-3">
                  <div className="flex items-center">
                    <Label htmlFor="password">Password</Label>
                  </div>
                    <Input id="password" type="password" required onChange={(e) => setPassword(e.target.value)}/>
                </div>
                <Button type="submit" className="w-full" onClick={() => login()}>
                  Login
                </Button>
              </div>
            </div>
        </CardContent>
      </Card>
    </div>
  )
}
