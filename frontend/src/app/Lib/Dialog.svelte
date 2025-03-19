<script>
    import BtnIcon from "./BtnIcon.svelte";
    import Icon from "./Icon.svelte";

    export let header = "";
    export let icon = {};
    export let btns = [];
    export let outside = false;
    export const mod = {
        show() {
            modal.style.display = "flex";
        },
        close() {
            modal.style.display = "none";
        },
    };

    let modal;

    function btnclick(fn) {
        mod.close();
        if (fn) {
            fn();
        }
    }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
    class="modal"
    bind:this={modal}
    on:click|self={() => {
        if (outside) {
            mod.close();
        }
    }}
>
    <div class="dialog">
        <div class="dialog-header">
            <div style="display:flex;align-items: center;padding-right: 3em;">
                {#if icon.name}
                    <Icon
                        name={icon.name}
                        color={icon.color ? icon.color : "var(--contrast)"}
                        size="2.25em"
                    />
                    <div style="width: 0.5em;" />
                {/if}
                <h4>{header}</h4>
            </div>
            <BtnIcon fn={mod.close} name="close" size="2em" />
        </div>
        <div
            class="dialog-content"
            class:bottom-radius={(!btns || btns.length == 0) && !$$slots.footer}
        >
            <slot name="content" />
        </div>
        {#if btns && btns.length > 0}
            <div class="dialog-footer">
                {#each btns as btn}
                    <button
                        style="min-width: 5em;margin-left: 1.5em"
                        class="btn {btn.class}"
                        on:click={btnclick(btn.callback)}
                    >
                        {btn.title}
                    </button>
                {/each}
            </div>
        {:else if $$slots.footer}
            <div class="dialog-footer">
                <slot name="footer" />
            </div>
        {/if}
    </div>
</div>
