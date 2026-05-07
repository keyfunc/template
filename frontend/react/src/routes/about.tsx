import { createFileRoute } from "@tanstack/react-router";
import About from "@/page/About";

export const Route = createFileRoute("/about")({
	component: About,
});
