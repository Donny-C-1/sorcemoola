<script>
	const pillars = $state([
		{
			title: "Our Mission",
			detail:
				"To empower Nigerians by providing the most secure and transparent crowdfunding platform dedicated to facilitating immediate critical aid and fostering sustainable personal advancement."
		},
		{
			title: "Our Vision",
			detail:
				"To fundamentally redefine the Nigerian social safety net, creating a future where every urgent need is met by the swift, collective power of the community, driven by integrity and trust."
		},
		{
			title: "Our Core Values",
			detail:
				"We operate on three non-negotiable pillars: Trust First (100% campaign verification), Rapid Impact (swift fund disbursement), and Communal Ownership (a platform built by and for Nigerians)."
		},
		{
			title: "Our Community",
			detail:
				"SorceMoola is a movement of millions of Nigerians at home and abroad, uniting to ensure that local challenges are met with global generosity, fostering resilience one successful campaign at a time."
		}
	]);

	// Team data including professional placeholder images
	const team = $state([
		{
			name: "Ayo Adekunle",
			title: "Founder & CEO",
			motivation:
				"Pioneering transparent giving and fighting financial vulnerability across Nigeria.",
			image: "https://i.pravatar.cc/300?img=70"
		},
		{
			name: "Nneka Okoro",
			title: "Head of Verification",
			motivation: "Ensuring ethical rigor and integrity is at the forefront of every campaign.",
			image: "https://i.pravatar.cc/300?img=47"
		},
		{
			name: "Tunde Bello",
			title: "Community Lead",
			motivation: "Empowering local campaigns and expanding SorceMoola's grassroots reach.",
			image: "https://i.pravatar.cc/300?img=33"
		},
		{
			name: "Femi Johnson",
			title: "Platform Architect",
			motivation: "Building the secure, fast, and scalable technology foundation for aid delivery.",
			image: "https://i.pravatar.cc/300?img=11"
		}
	]);

	const groupImage =
		"https://images.unsplash.com/photo-1552664730-d307ca884978?q=80&w=1470&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"; // Team collaboration photo

	const cta = $state({
		headline: "Join the Movement: Fund a Future, Start a Campaign.",
		subtext:
			"SorceMoola is the community you've been waiting for. Whether you're ready to give aid today or need a trusted platform to launch your personal advancement campaign, your journey starts here."
	});
</script>

<main>
	<section class="who_we_are">
		<div class="container">
			<div class="layout">
				<div>
					<div>
						<p class="subtitle">About Us</p>
						<h1 class="title">Who Exactly Are We?</h1>
						<p>
							SorceMoola is the secure, community-driven crowdfunding platform built to formalize
							and strengthen Nigeria’s innate culture of communal support. We exist to close the
							financial gap left by inadequate social safety nets, providing a transparent bridge
							between the desire to help and the critical needs on the ground.
						</p>
					</div>
					<img
						class="img"
						src="https://images.unsplash.com/photo-1518398046578-8cca57782e17?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHwxfHxnaXZpbmclMjB0byUyMG5lZWR5fGVufDB8fHx8MTc2NDgxMTY3N3ww&ixlib=rb-4.1.0&fit=max&q=80"
						alt=""
						width="100%"
						height="auto"
					/>
					<img
						class="img"
						src="https://images.unsplash.com/photo-1560220604-1985ebfe28b1?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHw1fHxibGFjayUyMHBlb3BsZSUyMHZvbHVudGVlcmluZ3xlbnwwfHx8fDE3NjQ4MTE5NDh8MA&ixlib=rb-4.1.0&fit=max&q=80"
						alt=""
						width="100%"
						height="auto"
					/>
				</div>
				<div>
					<img
						class="img"
						src="https://images.unsplash.com/photo-1575356891535-a79c0bc6bffc?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHw5fHxibGFjayUyMHBlb3BsZSUyMGNyb3dkZnVuZGluZ3xlbnwwfHx8fDE3NjQ4MTIwNjF8MA&ixlib=rb-4.1.0&fit=max&q=80"
						alt=""
						width="100%"
						height="auto"
					/>
					<div>
						<p>
							From providing essential resources to funding life-changing projects, every effort is
							directed toward building a better, more equitable world. By uniting individuals,
							businesses, and communities.
						</p>
						<ul class="checklist">
							<li>
								<i class="checkmark fa-solid fa-check"></i>Join Our Mission to Make a Difference
							</li>
							<li><i class="checkmark fa-solid fa-check"></i>Transforming Lives and Communities</li>
							<li><i class="checkmark fa-solid fa-check"></i>Standing Up for Human Rights</li>
						</ul>
					</div>
				</div>
			</div>
		</div>
	</section>

	<section class="mission_vision">
		<div class="container">
			<div class="quote-block">
				<blockquote class="quote-text">
					In Nigeria, family is the first safety net. SorceMoola is simply the digital extension of
					that love, ensuring help is verified, immediate, and trustworthy.
				</blockquote>
				<cite class="quote-attribution">— Chikwem Donald, Founder & CEO</cite>
			</div>

			<!-- 2. Core Pillars (Dual Column Layout) -->
			<div class="pillars-container">
				{#each pillars as pillar (pillar.title)}
					<div class="pillar-row">
						<!-- Left Column: Title -->
						<div class="pillar-title-col">
							<h3 class="pillar-title">{pillar.title}</h3>
						</div>

						<!-- Right Column: Detail Text -->
						<div class="pillar-detail-col">
							<p class="pillar-detail-text">{pillar.detail}</p>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</section>

	<div class="team-section">
		<div class="container">
			<!-- Section Title -->
			<h2 class="section-title">The Vanguards of Trust and Impact</h2>
			<p class="section-subtitle">
				Meet the leadership team committed to integrity, transparency, and the community.
			</p>

			<!-- 1. Group Visual (The Unity Element) -->
			<div class="group-visual-block">
				<img
					src={groupImage}
					alt="The SorceMoola leadership team collaborating"
					class="group-photo"
					loading="lazy"
				/>
				<div class="group-caption">
					<span class="caption-text"
						>"We build the platform, but the community is the true foundation."</span
					>
				</div>
			</div>

			<!-- 2. Individual Cards (The Accountability Element) -->
			<div class="team-grid">
				{#each team as member (member.name)}
					<div class="team-card">
						<div class="photo-wrapper">
							<img
								src={member.image}
								alt={`Portrait of ${member.name}, ${member.title}`}
								class="team-photo"
								loading="lazy"
							/>
						</div>
						<h3 class="team-name">{member.name}</h3>
						<p class="team-title">{member.title}</p>
						<p class="team-motivation">{member.motivation}</p>
					</div>
				{/each}
			</div>
		</div>
	</div>

	<section class="cta-section">
		<div class="container">
			<!-- Main CTA Card Block -->
			<div class="cta-card">
				<!-- Headline -->
				<h2 class="cta-headline">{cta.headline}</h2>

				<!-- Subtext -->
				<p class="cta-subtext">{cta.subtext}</p>

				<!-- Dual Call-to-Action Buttons -->
				<div class="cta-actions">
					<!-- Button A: Primary Action (Donate) -->
					<a href="/explore" class="btn btn-donate">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							stroke-width="2"
							stroke-linecap="round"
							stroke-linejoin="round"
							class="lucide lucide-heart-handshake"
							><path
								d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"
							/><path d="m12 14 6-6" /></svg
						>
						<span>Donate Now</span>
					</a>

					<!-- Button B: Secondary Action (Start Campaign) -->
					<a href="/campaign/start" class="btn btn-campaign">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							stroke-width="2"
							stroke-linecap="round"
							stroke-linejoin="round"
							class="lucide lucide-megaphone"
							><path d="m3 11 18-5v12L3 14v-3z" /><path d="M11.6 16.8a3 3 0 1 1-5.8-1.6" /><path
								d="M14 11h.01"
							/></svg
						>
						<span>Start a Campaign</span>
					</a>
				</div>
			</div>
		</div>
	</section>
</main>

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
		padding: 0.5rem;
		border-radius: 50%;
		font-size: 0.75rem;
		margin-right: var(--spacing-smr);
	}

	.quote-block {
		max-width: 55rem;
		margin: 0 auto var(--spacing-2xl);
		text-align: center;
		padding: var(--spacing-lg);
		border-bottom: 2px solid var(--primary-bg);
	}

	.quote-text {
		font-size: 1.75rem;
		font-style: italic;
		line-height: 1.4;
		color: var(--accent-color);
		margin: 0 0 var(--spacing-sm);
		padding: 0;
		position: relative;
	}

	.quote-text::before {
		content: "“";
		font-size: 4rem;
		position: absolute;
		top: -15px;
		left: -15px;
		opacity: 0.2;
		color: var(--accent-color);
		line-height: 1;
	}

	.quote-attribution {
		display: block;
		font-size: 1rem;
		font-weight: 600;
		color: var(--text-medium);
	}

	.pillars-container {
		border-top: 1px solid var(--primary-bg);
	}

	.pillar-row {
		display: grid;
		grid-template-columns: 1fr 2fr;
		gap: var(--spacing-lg);
		padding: var(--spacing-lg) 0;
		border-bottom: 1px solid var(--primary-bg);
		align-items: flex-start;
		transition: background-color 0.3s ease;
	}

	.pillar-title-col {
		padding-right: var(--spacing-md);
		padding-top: var(--spacing-xs);
	}

	.pillar-title {
		font-size: 1.5rem;
		font-weight: 700;
		color: var(--primary-dark);
		margin: 0;
	}

	.pillar-detail-col {
		padding-left: var(--spacing-md);
		border-left: 2px solid var(--primary-color);
	}

	.pillar-detail-text {
		font-size: 1.1rem;
		line-height: 1.7;
		color: var(--text-dark);
		margin: 0;
	}

	/* === RESPONSIVENESS === */
	@media (max-width: 1024px) {
		.quote-block {
			max-width: 100%;
		}

		.pillar-row {
			grid-template-columns: 1fr;
			gap: var(--spacing-sm);
			padding: var(--spacing-lg) 0;
		}

		.pillar-title-col {
			padding-right: 0;
		}

		.pillar-title {
			font-size: 1.35rem;
			color: var(--accent-color);
		}

		.pillar-detail-col {
			padding-left: var(--spacing-md);
		}
	}

	@media (max-width: 600px) {
		.quote-text {
			font-size: 1.4rem;
		}

		.quote-text::before {
			font-size: 3rem;
			top: -10px;
			left: 0;
		}

		.pillar-detail-text {
			font-size: 1rem;
		}
	}

	.team-section {
		background-color: var(--primary-bg); /* Subtle background to separate from white sections */
		padding: var(--spacing-3xl) 0;
		font-family: sans-serif;
	}

	.container {
		max-width: var(--container-width);
		margin: 0 auto;
		padding: 0 var(--container-padding);
	}

	.section-title {
		font-size: 2.5rem;
		font-weight: 700;
		color: var(--text-dark);
		margin-bottom: var(--spacing-xs);
		text-align: center;
	}

	.section-subtitle {
		font-size: 1.2rem;
		color: var(--text-medium);
		text-align: center;
		margin-bottom: var(--spacing-2xl);
	}

	/* === 1. Group Visual Block === */
	.group-visual-block {
		position: relative;
		max-width: 65rem;
		margin: 0 auto var(--spacing-2xl);
		border-radius: var(--radius-lg);
		overflow: hidden;
		box-shadow: var(--shadow-lg);
	}

	.group-photo {
		width: 100%;
		height: 350px;
		object-fit: cover;
		display: block;
		filter: brightness(0.7); /* Darken slightly for text overlay */
	}

	.group-caption {
		position: absolute;
		bottom: 0;
		left: 0;
		right: 0;
		padding: var(--spacing-md) var(--spacing-lg);
		background: linear-gradient(transparent, rgba(0, 0, 0, 0.7));
		text-align: center;
	}

	.caption-text {
		color: var(--text-white);
		font-size: 1.5rem;
		font-weight: 600;
		font-style: italic;
	}

	/* === 2. Individual Cards Grid === */
	.team-grid {
		display: grid;
		grid-template-columns: repeat(4, 1fr); /* 4 columns on desktop */
		gap: var(--spacing-lg);
		margin-top: var(--spacing-xl);
	}

	/* **UPDATED CARD STYLING** */
	.team-card {
		text-align: center;
		padding: var(--spacing-lg);
		background-color: var(--text-white);
		border-radius: var(--radius-md);
		/* Stronger default shadow */
		box-shadow: 0 4px 10px rgba(0, 0, 0, 0.08);
		transition:
			transform 0.3s ease,
			box-shadow 0.3s ease;
	}

	.team-card:hover {
		transform: translateY(-8px); /* Lifts the card more */
		/* More prominent hover shadow */
		box-shadow: 0 15px 30px rgba(0, 0, 0, 0.2);
	}

	/* **UPDATED PHOTO WRAPPER STYLING** */
	.photo-wrapper {
		width: 150px; /* Slightly larger image */
		height: 150px;
		margin: 0 auto var(--spacing-sm);
		border-radius: 50%;
		overflow: hidden;
		/* Dual ring border effect */
		border: 4px solid var(--primary-light);
		box-shadow: 0 0 0 2px var(--accent-color); /* Outer accent color ring */
		transition: all 0.3s ease;
	}

	.team-card:hover .photo-wrapper {
		box-shadow:
			0 0 0 6px var(--primary-light),
			0 0 0 8px var(--accent-color); /* Highlight border on hover */
	}

	.team-photo {
		width: 100%;
		height: 100%;
		object-fit: cover;
		display: block;
	}

	.team-name {
		font-size: 1.35rem; /* Slightly larger name */
		font-weight: 800; /* Bolder name */
		color: var(--text-dark);
		margin-bottom: var(--spacing-xs);
	}

	.team-title {
		/* Using accent color for the role title */
		color: var(--accent-color);
		font-weight: 600;
		font-size: 1rem;
		margin-bottom: var(--spacing-sm);
	}

	.team-motivation {
		font-size: 0.95rem;
		color: var(--text-medium);
		line-height: 1.4;
	}

	/* === RESPONSIVENESS === */
	@media (max-width: 1200px) {
		.team-grid {
			grid-template-columns: repeat(2, 1fr); /* 2 columns on medium screens */
		}
		.group-photo {
			height: 300px;
		}
	}

	@media (max-width: 600px) {
		.section-title {
			font-size: 2rem;
		}

		.section-subtitle {
			font-size: 1rem;
		}

		.team-grid {
			grid-template-columns: 1fr; /* 1 column on mobile */
		}

		.group-photo {
			height: 200px;
		}

		.caption-text {
			font-size: 1.2rem;
		}

		.photo-wrapper {
			width: 120px;
			height: 120px;
		}
	}

    .cta-section {
        /* This section should stand out, hence a slight color change */
        background-color: var(--text-white); 
        padding: var(--spacing-3xl) 0;
        font-family: sans-serif;
    }
    
    .container {
        max-width: var(--container-width);
        margin: 0 auto;
        padding: 0 var(--container-padding);
    }

    /* === CTA Card Styling === */
    .cta-card {
        background-color: var(--secondary-color); /* Bold Secondary Color */
        color: var(--text-dark);
        padding: var(--spacing-2xl);
        border-radius: var(--radius-lg);
        text-align: center;
        max-width: 60rem;
        margin: 0 auto;
        /* Prominent Shadow to stand out */
        box-shadow: 0 20px 40px rgba(248, 211, 71, 0.6); 
    }

    .cta-headline {
        font-size: 2.5rem;
        font-weight: 800;
        line-height: 1.2;
        color: var(--text-dark);
        margin-bottom: var(--spacing-md);
    }

    .cta-subtext {
        font-size: 1.15rem;
        line-height: 1.6;
        color: var(--text-medium);
        max-width: 45rem;
        margin: 0 auto var(--spacing-xl);
    }

    /* === Dual Actions/Buttons === */
    .cta-actions {
        display: flex;
        justify-content: center;
        gap: var(--spacing-lg);
    }

    .btn {
        display: inline-flex;
        align-items: center;
        gap: var(--spacing-sm);
        padding: 1rem 1.75rem;
        border: none;
        border-radius: var(--radius-md);
        text-decoration: none;
        font-weight: 700;
        font-size: 1.1rem;
        transition: transform 0.2s ease, box-shadow 0.3s ease;
        cursor: pointer;
        box-shadow: var(--shadow-md);
    }

    .btn:hover {
        transform: translateY(-2px);
        box-shadow: var(--shadow-lg);
    }

    /* Primary CTA: Donate (Green Theme) */
    .btn-donate {
        background-color: var(--primary-color);
        color: var(--text-white);
    }
    .btn-donate:hover {
        background-color: var(--primary-dark);
    }
    
    /* Secondary CTA: Campaign (Accent Theme - Purple) */
    .btn-campaign {
        background-color: var(--accent-color);
        color: var(--text-white);
    }
    .btn-campaign:hover {
        background-color: var(--accent-dark);
    }
    
    .btn svg {
        width: 20px;
        height: 20px;
    }

    /* === RESPONSIVENESS === */
    @media (max-width: 768px) {
        .cta-card {
            padding: var(--spacing-lg);
        }
        
        .cta-headline {
            font-size: 2rem;
        }
        
        .cta-subtext {
            font-size: 1rem;
            margin-bottom: var(--spacing-lg);
        }
        
        .cta-actions {
            flex-direction: column;
            gap: var(--spacing-sm);
        }
        
        .btn {
            width: 100%;
            padding: 0.8rem 1.5rem;
            font-size: 1rem;
            justify-content: center;
        }
    }
    
    @media (max-width: 480px) {
         .cta-headline {
            font-size: 1.75rem;
        }
    }
</style>
