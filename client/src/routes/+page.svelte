<script>
    // This is a minimal Svelte 5 script block.
    // We use $state() only to demonstrate the Svelte 5 runic syntax,
    // though this static page requires no complex state management.
    let headline = $state('Collect Funds. No Hassle.');
    let primaryColor = '#10B981'; // Emerald Green
</script>

<svelte:head>
    <!-- Note: In a real Svelte project, Tailwind CSS is imported or built. 
         Here, we assume the classes are available for styling. -->
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800&display=swap" rel="stylesheet">
    <title>Sorcemoola - Fund Anything, Effortlessly</title>
</svelte:head>

<main class="min-h-screen bg-gray-50 text-gray-900 font-inter">
    <!-- Navigation -->
    <header class="py-4 border-b border-gray-200">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex justify-between items-center">
            <a href="./#" class="text-2xl font-extrabold tracking-tight">
                Sorcemoola<span style="color: {primaryColor};">.</span>
            </a>
            <a href="#cta" class="bg-white text-gray-900 font-semibold py-2 px-4 rounded-full border border-gray-300 hover:bg-gray-100 transition duration-300 shadow-sm">
                Log In
            </a>
        </div>
    </header>

    <!-- Hero Section -->
    <section class="text-center py-20 md:py-32 overflow-hidden">
        <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
            <h1 class="text-5xl md:text-7xl font-extrabold leading-tight mb-6">
                <span class="text-gray-900">{@html headline.split('. ')[0]}</span> 
                <span style="color: {primaryColor};" class="block sm:inline">{@html headline.split('. ')[1]}</span>
            </h1>
            <p class="text-xl md:text-2xl text-gray-600 mb-10 max-w-2xl mx-auto">
                The universal platform for students, events, and any purpose. Stop chasing account numbers and start collecting instantly.
            </p>
            <!-- Primary CTA -->
            <a href="#how-it-works" id="cta" 
               class="inline-block text-white font-bold text-lg py-3 px-10 rounded-full transition duration-300 transform scale-105 shadow-xl"
               style="background-color: {primaryColor}; box-shadow: 0 4px 15px rgba(16, 185, 129, 0.4);"
               onmouseenter={this.style.boxShadow='0 8px 20px rgba(16, 185, 129, 0.6)'}
               onmouseleave={this.style.boxShadow='0 4px 15px rgba(16, 185, 129, 0.4)'}
            >
                Start Your Fundraiser Today
            </a>
            <p class="mt-4 text-sm text-gray-500">Lowest fees guaranteed. Seriously simple setup.</p>
        </div>
    </section>

    <!-- Value Proposition / Differentiators -->
    <section class="py-16 md:py-24 bg-white">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <h2 class="text-4xl font-bold text-center mb-16 text-gray-900">Why Sorcemoola is Better</h2>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-8">

                <!-- Card 1: Instant Collections -->
                <div class="bg-gray-100 p-8 rounded-xl shadow-lg transition duration-500 hover:scale-[1.02] border border-gray-200">
                    <div class="text-4xl mb-4" style="color: {primaryColor};">✨</div>
                    <h3 class="text-2xl font-semibold mb-3 text-gray-900">Instant Collections</h3>
                    <p class="text-gray-600">Share one simple, personalized link. Friends, family, or your community can contribute securely in seconds—no more back-and-forth chats for bank details.</p>
                </div>

                <!-- Card 2: Low, Transparent Fees -->
                <div class="bg-gray-100 p-8 rounded-xl shadow-lg transition duration-500 hover:scale-[1.02] border border-gray-200">
                    <div class="text-4xl mb-4" style="color: {primaryColor};">💰</div>
                    <h3 class="text-2xl font-semibold mb-3 text-gray-900">Low, Transparent Fees</h3>
                    <p class="text-gray-600">We're built to keep more money in your pocket. Our fees are among the lowest in the industry, clearly stated, and never hidden.</p>
                </div>

                <!-- Card 3: For Every Purpose -->
                <div class="bg-gray-100 p-8 rounded-xl shadow-lg transition duration-500 hover:scale-[1.02] border border-gray-200">
                    <div class="text-4xl mb-4" style="color: {primaryColor};">🎯</div>
                    <h3 class="text-2xl font-semibold mb-3 text-gray-900">Fund Anything, Effortlessly</h3>
                    <p class="text-gray-600">Running a club event? Need to fund new school equipment? Or maybe a personal goal? Sorcemoola handles all your versatile fundraising needs seamlessly.</p>
                </div>
            </div>
        </div>
    </section>

    <!-- How It Works (3 Steps) -->
    <section id="how-it-works" class="py-16 md:py-24">
        <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
            <h2 class="text-4xl font-bold text-center mb-16 text-gray-900">How It Works in 3 Simple Steps</h2>
            <div class="relative">
                <!-- Mobile & Desktop Separator Line -->
                <div class="absolute hidden md:block w-0.5 h-full bg-gray-300 left-1/2 transform -translate-x-1/2"></div>
                <div class="space-y-12">
                    
                    <!-- Step 1 (Left on Desktop) -->
                    <div class="flex flex-col md:flex-row items-center relative">
                        <div class="md:w-1/2 md:pr-12 text-center md:text-right">
                            <h3 class="text-3xl font-bold mb-2" style="color: {primaryColor};">1. Create Your Campaign</h3>
                            <p class="text-gray-700 text-lg">Quickly define your goal, add a photo, and set your target. It takes less than 5 minutes to launch your page.</p>
                        </div>
                        <div class="md:w-1/2 flex justify-center mt-6 md:mt-0">
                             <div class="w-12 h-12 rounded-full text-white font-bold flex items-center justify-center text-xl absolute md:static border-4 border-gray-50 md:order-2 shadow-md" style="background-color: {primaryColor};">1</div>
                        </div>
                    </div>

                    <!-- Step 2 (Right on Desktop) -->
                    <div class="flex flex-col md:flex-row items-center relative md:pt-12">
                        <div class="md:w-1/2 md:pl-12 md:order-2 text-center md:text-left">
                            <h3 class="text-3xl font-bold mb-2" style="color: {primaryColor};">2. Share Your Unique Link</h3>
                            <p class="text-gray-700 text-lg">Share your Sorcemoola link on social media, in group chats, or via email. All payments are secured and tracked automatically.</p>
                        </div>
                        <div class="md:w-1/2 flex justify-center mt-6 md:mt-0 md:order-1">
                            <div class="w-12 h-12 rounded-full text-white font-bold flex items-center justify-center text-xl absolute md:static border-4 border-gray-50 shadow-md" style="background-color: {primaryColor};">2</div>
                        </div>
                    </div>

                    <!-- Step 3 (Left on Desktop) -->
                    <div class="flex flex-col md:flex-row items-center relative md:pt-12">
                        <div class="md:w-1/2 md:pr-12 text-center md:text-right">
                            <h3 class="text-3xl font-bold mb-2" style="color: {primaryColor};">3. Withdraw and Celebrate</h3>
                            <p class="text-gray-700 text-lg">Watch your funds grow in real-time. When you're ready, simply withdraw the collected amount directly to your bank account.</p>
                        </div>
                        <div class="md:w-1/2 flex justify-center mt-6 md:mt-0">
                            <div class="w-12 h-12 rounded-full text-white font-bold flex items-center justify-center text-xl absolute md:static border-4 border-gray-50 md:order-2 shadow-md" style="background-color: {primaryColor};">3</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <!-- Secondary CTA Strip -->
    <section class="py-12 md:py-16 bg-emerald-50 border-t border-b border-gray-200">
        <div class="max-w-4xl mx-auto text-center px-4 sm:px-6">
            <h2 class="text-4xl font-extrabold text-gray-900 mb-4">Ready to Moola?</h2>
            <p class="text-xl text-gray-700 mb-8">Join the movement making fundraising easier for millions of people, starting today.</p>
            <a href="#how-it-works" 
               class="inline-block text-white font-bold text-lg py-3 px-12 rounded-full transition duration-300 transform shadow-xl"
               style="background-color: {primaryColor};"
               onmouseenter={this.style.backgroundColor='#059669'}
               onmouseleave={this.style.backgroundColor='{primaryColor}'}
            >
                Launch My First Campaign
            </a>
        </div>
    </section>

    <!-- Footer -->
    <footer class="py-8 bg-white border-t border-gray-200">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center text-gray-500">
            <p class="mb-4">&copy; 2025 Sorcemoola. All rights reserved. | Simple Funding, Maximum Impact.</p>
            <div class="flex justify-center space-x-6 text-sm">
                <a href="./#" class="hover:text-emerald-500 transition">Terms of Service</a>
                <a href="./#" class="hover:text-emerald-500 transition">Privacy Policy</a>
                <a href="./#" class="hover:text-emerald-500 transition">Contact Us</a>
            </div>
        </div>
    </footer>
</main>