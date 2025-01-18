<script lang="ts">

    import type { PageData } from "../../../.svelte-kit/types/src/routes/upload/$types.js";

    let { data }: {data: PageData} = $props();

    let input;
    let image = $state();
    let showImage = $state(false);

    function onChange() {
        const file = input.files[0];

        if (file) {
            showImage = true;

            const reader = new FileReader();
            reader.addEventListener("load", function () {
                image.setAttribute("src", reader.result);
            });
            reader.readAsDataURL(file);

            return;
        }
        showImage = false;
    }

</script>


<div class="flex flex-col bg-gray-950 justify-center items-center w-screen h-screen text-white">
    <form method="POST" enctype=multipart/form-data action="?/upload" class="flex flex-col justify-center items-center">
        <h1 class="flex justify-center font-bold">Upload your media here!</h1>
        <div class="flex flex-col w-72 whitespace-nowrap overflow-hidden cursor-pointer">
            <input
                bind:this={input}
                onchange={onChange}
                type="file"
                class="text-sm file:rounded-lg file:border-0 file:p-2 m-10 text-ellipsis overflow-hidden w-full h-full break-words cursor-pointer"
            />
        </div>
        <div class="flex justify-center font-bold">
            {#if showImage}
                <img bind:this={image} src="" alt="Preview" class="w-64 h-auto aspect-square object-cover" />
            {:else}
                <span class="flex p-20 border-2">Image Preview</span>
            {/if}
        </div>

        <button class="flex justify-center items-center bg-gray-700 p-4 rounded-2xl w-fit m-10 font-bold text-gray-300 hover:scale-105 hover:bg-gray-500 transition-all">Finish and Upload</button>
    </form>
</div>