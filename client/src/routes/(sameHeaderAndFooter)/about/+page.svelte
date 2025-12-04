<script>
    const content = $state({
        subtitle: "About Us",
        title: "Who Exactly Are We?",
        missionText: "SorceMoola is the secure, community-driven crowdfunding platform built to formalize and strengthen Nigeria’s innate culture of communal support. We exist to close the financial gap left by inadequate social safety nets, providing a transparent bridge between the desire to help and the critical needs on the ground.",
        focusTitle: "Our Unwavering Focus on Trust and Impact",
        focusText: "We are exclusively focused on funding Critical Aid (medical, education, relief) and Personal Advancement (micro-business, skills acquisition) for verified Nigerians. This specialized approach ensures that every contribution is protected by our rigorous verification process, translating directly into life-changing support, swiftly and securely."
    });
</script>

<section class="who_we_are">
    <div class="container">
        <div class="layout">
            <div>
                <div>
                    <p class="subtitle">About Us</p>
                    <h1 class="title">Who Exactly Are We?</h1>
                    <p>SorceMoola is the secure, community-driven crowdfunding platform built to formalize and strengthen Nigeria’s innate culture of communal support. We exist to close the financial gap left by inadequate social safety nets, providing a transparent bridge between the desire to help and the critical needs on the ground.</p>
                </div>
                <img class="img" src="https://images.unsplash.com/photo-1518398046578-8cca57782e17?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHwxfHxnaXZpbmclMjB0byUyMG5lZWR5fGVufDB8fHx8MTc2NDgxMTY3N3ww&ixlib=rb-4.1.0&fit=max&q=80" alt="" width="100%" height="auto" />
                <img class="img" src="https://images.unsplash.com/photo-1560220604-1985ebfe28b1?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHw1fHxibGFjayUyMHBlb3BsZSUyMHZvbHVudGVlcmluZ3xlbnwwfHx8fDE3NjQ4MTE5NDh8MA&ixlib=rb-4.1.0&fit=max&q=80" alt="" width="100%" height="auto" />
            </div>
            <div>
                <img class="img" src="https://images.unsplash.com/photo-1575356891535-a79c0bc6bffc?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHw5fHxibGFjayUyMHBlb3BsZSUyMGNyb3dkZnVuZGluZ3xlbnwwfHx8fDE3NjQ4MTIwNjF8MA&ixlib=rb-4.1.0&fit=max&q=80" alt="" width="100%" height="auto" />
                <div>
                    <p>From providing essential resources to funding life-changing projects, every effort is directed toward building a better, more equitable world. By uniting individuals, businesses, and communities.</p>
                    <ul class="checklist">
                        <li><i class="checkmark fa-solid fa-check"></i>Join Our Mission to Make a Difference</li>
                        <li><i class="checkmark fa-solid fa-check"></i>Transforming Lives and Communities</li>
                        <li><i class="checkmark fa-solid fa-check"></i>Standing Up for Human Rights</li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</section>


<style>
    .who_we_are .layout {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 2rem;
    }

    @media (max-width: 47rem) {
        .who_we_are .layout {
            grid-template-columns: 1fr;
        }
    }

    .who_we_are .layout > div {
        display: grid;
        gap: 2rem;
    }

    .subtitle {
        display: inline-block;
        font-size: 0.9rem;
        font-weight: 600;
        text-transform: uppercase;
        padding: var(--spacing-xs) var(--spacing-sm);
        border-radius: var(--radius-sm);
        border: 1px solid var(--neutral-color);
        margin-bottom: var(--spacing-sm);
    }

    .title {
        font-size: 3rem;
        font-weight: 900;
        line-height: 1.1;
        margin-bottom: var(--spacing-md);
    }

    .checklist {
        list-style-type: none;
        padding: 0;
    }

    .checklist li {
        margin-bottom: var(--spacing-sm);
        font-weight: 700;
    }

    .checkmark {
        background-color: var(--secondary-color);
        padding: .5rem;
        border-radius: 50%;
        font-size: .75rem;
        margin-right: var(--spacing-smr);
    }
</style>