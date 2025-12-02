<footer>
	<div class="container">
		<div class="footer-grid">
			<div class="footer-about">
				<div class="footer-logo">
					<span class="logo-text">Sorce<span class="highlight">Moola</span></span>
				</div>
				<p class="tagline">Empowering innovators. Funding the future.</p>
				<div class="social-icons">
					<a href="./#" aria-label="Facebook"><i class="fab fa-facebook-f"></i></a>
					<a href="./#" aria-label="Twitter"><i class="fab fa-twitter"></i></a>
					<a href="./#" aria-label="Instagram"><i class="fab fa-instagram"></i></a>
					<a href="./#" aria-label="LinkedIn"><i class="fab fa-linkedin-in"></i></a>
				</div>
			</div>

			<div class="footer-links">
				<h3>Quick Links</h3>
				<ul>
					<li><a href="/about">About Us</a></li>
					<li><a href="/terms">Terms of Service</a></li>
					<li><a href="/privacy">Privacy Policy</a></li>
					<li><a href="/help">Help & Support</a></li>
				</ul>
			</div>

			<div class="footer-newsletter">
				<h3>Stay Updated</h3>
				<p>Subscribe to our newsletter for the latest campaigns and updates</p>
				<div class="newsletter-form">
					<input type="email" placeholder="Your email address" />
					<button type="submit">Subscribe</button>
				</div>
			</div>
		</div>

		<div class="footer-bottom">
			<p>© {new Date().getFullYear()} SorceMoola. All rights reserved.</p>
		</div>
	</div>
</footer>

<style>
	footer {
		background: #272a33;
		color: white;
		padding: 4rem 0 1.5rem;
	}

	.footer-grid {
		display: grid;
		grid-template-columns: 2fr 1fr 2fr;
		gap: 3rem;
		margin-bottom: 3rem;
	}

	.footer-about {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.footer-logo {
		font-size: 1.5rem;
		font-weight: 700;
	}

	.tagline {
		color: #ccc;
	}

	.social-icons {
		display: flex;
		gap: 1rem;
		margin-top: 1rem;
	}

	.social-icons a {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 2.25rem;
		height: 2.25rem;
		border-radius: 50%;
		background: rgba(255, 255, 255, 0.1);
		transition: background 0.3s ease;
	}

	.social-icons a:hover {
		background: #3cb983;
	}

	.footer-links h3,
	.footer-newsletter h3 {
		font-size: 1.2rem;
		margin-bottom: 1.5rem;
	}

	.footer-links ul {
		list-style: none;
		padding: 0;
		margin: 0;
	}

	.footer-links li {
		margin-bottom: 0.75rem;
	}

	.footer-links a {
		color: #ccc;
		transition: color 0.3s ease;
	}

	.footer-links a:hover {
		color: #3cb983;
	}

	.newsletter-form {
		display: flex;
		margin-top: 1rem;
	}

	.newsletter-form input {
		flex-grow: 1;
		padding: 0.75rem 1rem;
		border: none;
		border-radius: 4px 0 0 4px;
		outline: none;
		width: 100%;
	}

	.newsletter-form button {
		background: #3cb983;
		color: white;
		border: none;
		padding: 0.75rem 1.25rem;
		border-radius: 0 4px 4px 0;
		font-weight: 500;
		transition: background 0.3s ease;
	}

	.newsletter-form button:hover {
		background: #34a273;
	}

	.footer-bottom {
		text-align: center;
		padding-top: 1.5rem;
		border-top: 1px solid rgba(255, 255, 255, 0.1);
		color: #aaa;
		font-size: 0.9rem;
	}

    @media (max-width: 768px) {
		.footer-grid {
			grid-template-columns: 1fr;
			gap: 2rem;
		}

		.social-icons {
			justify-content: center;
		}

		.footer-links {
			text-align: center;
		}

		.footer-links ul {
			display: flex;
			flex-wrap: wrap;
			justify-content: center;
			gap: 1rem;
		}

		.footer-newsletter {
			text-align: center;
		}
	}
</style>