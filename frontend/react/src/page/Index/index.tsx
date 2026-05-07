import type { TodoPageResponse } from "@/service/generated/model";
import { useListTodos } from "@/service/generated/todo/todo";

const isTodoPageResponse = (
	response: TodoPageResponse | { data: unknown } | undefined,
): response is TodoPageResponse => {
	return (
		typeof response?.data === "object" &&
		response.data !== null &&
		"list" in response.data
	);
};

function Index() {
	const todoQuery = useListTodos(
		{ page: 1, size: 10 },
		{
			query: {
				refetchOnWindowFocus: false,
			},
		},
	);

	const todoPage = isTodoPageResponse(todoQuery.data)
		? todoQuery.data
		: undefined;
	const todos = todoPage?.data.list ?? [];

	return (
		<main className="mx-auto flex w-full max-w-4xl flex-col gap-6 px-6 py-8">
			<header className="flex flex-wrap items-center justify-between gap-3">
				<div>
					<h1 className="text-2xl font-semibold text-zinc-950">Todo List</h1>
					<p className="mt-1 text-sm text-zinc-500">
						共 {todoPage?.data.total ?? 0} 条，当前第 {todoPage?.data.page ?? 1}{" "}
						页
					</p>
				</div>
				<button
					className="rounded-md bg-zinc-950 px-4 py-2 text-sm font-medium text-white disabled:cursor-not-allowed disabled:bg-zinc-300"
					disabled={todoQuery.isFetching}
					type="button"
					onClick={() => {
						void todoQuery.refetch();
					}}
				>
					{todoQuery.isFetching ? "刷新中" : "刷新"}
				</button>
			</header>

			<section className="overflow-hidden rounded-lg border border-zinc-200 bg-white">
				{todoQuery.isLoading ? (
					<div className="px-5 py-8 text-sm text-zinc-500">正在获取 Todo</div>
				) : null}

				{todoQuery.isError ? (
					<div className="px-5 py-8 text-sm text-red-600">
						{todoQuery.error.message}
					</div>
				) : null}

				{!todoQuery.isLoading && !todoQuery.isError && todos.length === 0 ? (
					<div className="px-5 py-8 text-sm text-zinc-500">暂无 Todo</div>
				) : null}

				{todos.length > 0 ? (
					<ul className="divide-y divide-zinc-100">
						{todos.map((todo) => (
							<li className="flex flex-col gap-3 px-5 py-4" key={todo.id}>
								<div className="flex flex-wrap items-center justify-between gap-2">
									<h2 className="text-base font-medium text-zinc-950">
										{todo.title}
									</h2>
									<span className="rounded-md bg-zinc-100 px-2 py-1 text-xs font-medium text-zinc-600">
										状态 {todo.status}
									</span>
								</div>
								<p className="text-sm leading-6 text-zinc-600">
									{todo.description}
								</p>
								<div className="flex flex-wrap gap-4 text-xs text-zinc-400">
									<span>ID: {todo.id}</span>
									<span>创建: {todo.create_at}</span>
									<span>更新: {todo.update_at}</span>
								</div>
							</li>
						))}
					</ul>
				) : null}
			</section>
		</main>
	);
}

export default Index;
