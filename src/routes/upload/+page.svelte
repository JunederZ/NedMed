<script lang="ts">

    import type {ActionData} from "../../../.svelte-kit/types/src/routes/upload/$types.js";
    import { enhance } from '$app/forms';

    let { form }: { form: ActionData } = $props();

    let input: HTMLInputElement;
    let image = $state<HTMLImageElement>();
    let showImage = $state(false);

    function onChange() {
        const file = input.files?.[0];

        if (file) {
            showImage = true;

            const reader = new FileReader();
            reader.addEventListener("load", function () {
                if (image && reader.result != null && typeof reader.result === "string") {
                    image.setAttribute("src", reader.result);
                }
            });
            reader.readAsDataURL(file);

            return;
        }
        showImage = false;
    }

</script>


<div class="flex flex-col bg-gray-950 justify-center items-center w-screen h-screen text-white">
    <form
            method="POST"
            enctype=multipart/form-data
            class="flex flex-col justify-center items-center"
            use:enhance={() => {
                return async ({update}) => {
                    showImage = false
                    await update()
                }
            }}
    >
        <h1 class="flex justify-center font-bold">Upload your media here!</h1>
        <div class="flex flex-col w-72 whitespace-nowrap overflow-hidden">
            <input
                bind:this={input}
                onchange={onChange}
                type="file"
                required
                id="file"
                accept="image/*"
                name="file"
                class="text-sm file:rounded-lg file:border-0 file:p-2 m-10 text-ellipsis w-full cursor-pointer"
            />
            <label for="title">
                Image title
            </label>
            <input
                    required
                    type="text"
                    id="title"
                    name="title"
                    placeholder="Title"
                    class="text-gray-600 p-2 focus:outline-none my-2 rounded-lg bg-gray-100"
            >
            <label for="description">
                Image description
            </label>
            <textarea
                    required
                    id="description"
                    name="description"
                    placeholder="Description"
                    class="text-gray-600 text-wrap p-2 focus:outline-none my-2 rounded-lg bg-gray-100 max-h-24"
            ></textarea>
        </div>
        <div class="flex justify-center font-bold">
            {#if showImage}
                <img bind:this={image} src="" alt="Preview" class="w-64 h-auto aspect-square object-cover" />
            {:else}
                <span class="flex p-20 border-2">Image Preview</span>
            {/if}
        </div>

        <button class="flex justify-center items-center bg-gray-700 p-4 rounded-2xl w-fit m-10 font-bold text-gray-300 hover:scale-105 hover:bg-gray-500 transition-all">Finish and Upload</button>
        {#if form?.success}
            <p>Image uploaded successfully!</p>
        {/if}
    </form>
</div>