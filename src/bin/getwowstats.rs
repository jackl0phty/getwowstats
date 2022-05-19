// Import the getwowstats library to use in our CLI tool.
use clap::Parser;
use std::collections::HashMap;
use std::env;

#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]

struct Args {
    /// World of Warcraft character name
    #[clap(short, long)]
    character: String,

    /// Your two letter country code ( e.g. locale )
    #[clap(short, long)]
    locale: String,

    /// Method in getwowstats to call
    #[clap(short, long)]
    method: String,

    /// Realm of WoW character to lookup stats about
    #[clap(short, long)]
    //counter: u8,
    realm: String,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args = Args::parse();
    
    println!("Character {}", args.character);
    println!("Locale {}", args.locale);
    println!("Method {}", args.method);
    println!("Realm {}", args.realm);

    // Declare variables
    //base_api_url = "https://us.api.blizzard.com/profile/wow/character/";
    //forward_slash = "/";
    //let wow_api_endpoint = format!("{}{}{}{}{}{}{}", base_api_url, args.realm, forward_slash, args.character, "?namespace=profile-", args.region, "&locale=en_US&access_token=");
    
    let wow_api_token = "WOW_API_TOKEN";
    let wow_token = env::var("WOW_API_TOKEN").unwrap_or("none".to_string());
    println!("WOW_API_TOKEN is set to {}", wow_api_token);

    // Declare local variables
    //let wow_api_token = "abc123";
    let base_api_url = "https://us.api.blizzard.com/profile/wow/character/".to_owned();
    let base_api_url2 = base_api_url + &args.realm.to_owned();
    let base_api_url3 = base_api_url2 + &"/".to_owned();
    let base_api_url4 = base_api_url3 + &args.character.to_owned();
    let base_api_url5 = base_api_url4 + &"?namespace=profile-".to_owned();
    let base_api_url6 = base_api_url5 + &args.locale.to_owned();
    let base_api_url7 = base_api_url6 + &"&locale=en_US&access_token=".to_owned();
    let wow_api_url = base_api_url7 + &wow_token.to_owned();

    println!("{}", wow_api_url);

    //let resp = reqwest::get("https://httpbin.org/ip")
    let resp = reqwest::get(wow_api_url)
        .await?
        .json::<HashMap<String, String>>()
        .await?;
    println!("{:#?}", resp);
    Ok(())

}
