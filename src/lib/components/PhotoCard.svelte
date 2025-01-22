<script lang="ts">

    import NoImg from "../assets/no-img.png"

    let props = $props();
    
    const downloadImage = () => {
        if (props.url) {
            const link = document.createElement('a');
            link.href = props.url;
            link.download = props.title || 'downloaded_image';
            link.style.display = 'none'; 
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        }
    };

</script>


<div class="flex flex-col h-full bg-gray-600/60 p-4 rounded-xl justify-between items-center w-64">
    <div class="flex flex-col justify-center items-center">
        <div class="flex flex-col p-2 font-semibold max-w-32 text-center justify-center items-center">
            {#if props.filename == null || props.filename === ""}
                <h1 class="max-w-full text-center">No Title</h1>
            {:else}
                <div class="relative group">
                    <h1 class="line-clamp-1 max-w-full text-center cursor-help">
                        {props.filename}
                    </h1>
                    <div class="absolute hidden group-hover:block bg-gray-800 text-white p-2 rounded-md -top-8 left-1/2 transform -translate-x-1/2 max-w-48 break-words text-sm">
                        {props.filename}
                    </div>
                </div>
            {/if}
        </div>
        <div class="flex">
            {#if props.url.includes("example")}
                <img src={NoImg} alt="" class="w-auto h-40 aspect-square object-cover rounded-md" />
            {:else}
                <img src={props.img} alt="" class="w-auto h-40 aspect-square object-cover rounded-md" />
            {/if}
        </div>
    </div>
    <div class="flex flex-col justify-between gap-4 max-w-full p-2 h-36">
        {#if props.description == null || props.description === ""}
            <p>No Description.</p>
        {:else}
            <div class="relative group">
                <p class="justify-self-start line-clamp-2 cursor-help break-words w-full">
                    {props.description}
                </p>
                <div class="absolute hidden group-hover:block bg-gray-800 text-white p-2 rounded-md -top-8 left-1/2 -translate-x-1/2 z-10 min-w-44 max-w-96 break-words text-sm">
                    {props.description}
                </div>
            </div>
        {/if}
        <div>
            <button class="bg-gray-800 hover:bg-gray-900 hover:scale-105 py-2 px-4 my-2 rounded-3xl transition-all" onclick={downloadImage} disabled={!props.url}>Download</button>
        </div>
    </div>
</div>