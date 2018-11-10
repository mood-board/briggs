import React from 'react';



const BrowseByCategory = () => {
	return (
		<section className="py-8">
			<h2 className="text-center">Browse By Category</h2>
			<div className="font-sans flex justify-between my-4 mx-auto container overflow-hidden">
				<div class="cats pure-g hover_links cat_listings">
					<h3>
						<a href="/en/photos/?cat=animals">Animals</a></h3>
					<h3>
						<a href="/en/photos/?cat=buildings">Architecture/Buildings</a></h3>
					<h3>
						<a href="/en/photos/?cat=backgrounds">Backgrounds/Textures</a></h3>
					<h3>
						<a href="/en/photos/?cat=fashion">Beauty/Fashion</a></h3>
					<h3>
						<a href="/en/photos/?cat=business">Business/Finance</a></h3>
					<h3>
						<a href="/en/photos/?cat=computer">Computer/Communication</a></h3>
					<h3>
						<a href="/en/photos/?cat=education">Education</a></h3>
					<h3>
						<a href="/en/photos/?cat=feelings">Emotions</a></h3>
					<h3>
						<a href="/en/photos/?cat=food">Food/Drink</a></h3>
					<h3>
						<a href="/en/photos/?cat=health">Health/Medical</a></h3>
					<h3>
						<a href="/en/photos/?cat=industry">Industry/Craft</a></h3>
					<h3>
						<a href="/en/photos/?cat=music">Music</a></h3>
					<h3>
						<a href="/en/photos/?cat=nature">Nature/Landscapes</a></h3>
					<h3>
						<a href="/en/photos/?cat=people">People</a></h3>
					<h3>
						<a href="/en/photos/?cat=places">Places/Monuments</a></h3>
					<h3>
						<a href="/en/photos/?cat=religion">Religion</a></h3>
					<h3>
						<a href="/en/photos/?cat=science">Science/Technology</a></h3>
					<h3>
						<a href="/en/photos/?cat=sports">Sports</a></h3>
					<h3>
						<a href="/en/photos/?cat=transportation">Transportation/Traffic</a></h3>
					<h3>
						<a href="/en/photos/?cat=travel">Travel/Vacation</a></h3>
				</div>
			</div>
		</section>
	)
}

export default BrowseByCategory