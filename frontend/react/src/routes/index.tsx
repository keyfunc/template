import { createFileRoute } from "@tanstack/react-router";
import Index from "@/page/Index";

export const Route = createFileRoute("/")({
	component: Index,
});
