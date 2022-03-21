import { serve } from "https://deno.land/std@0.130.0/http/server.ts";

const port = 8080;
console.log(`HTTP server listening on http://localhost:${port}`);

serve(async (request: Request) => {
  const data = await request.json()
  return new Response(JSON.stringify({ data }));
}, {port});
