// Generated with prosr1 by compiler 0.9, template 0.1

using Microsoft.AspNetCore.SignalR;
using Microsoft.AspNetCore.SignalR.Client;
using System;
using System.Threading.Tasks;

namespace Test.TestNamespace
{
	public interface ISearchHubClient : ISearchHubClientBase, IDisposable
	{
		Task InitializeHubConnection();
		Task CallRefreshSearchTreeOnHub();
		Task CallTriggerOnHub();
		Task CallSearchOnHub(SearchRequest message);
	}

	public interface ISearchHubClientBase
	{
		Task SearchCalledOnAllAsync(SearchRequest message);
		Task SearchTreeRefreshedOnCallerAsync();
		Task SearchOnCallerAsync(SearchResponse message);
	}

	public abstract class SearchHubClientBase : ISearchHubClient
	{
		private readonly IHubConnectionBuilder _hubConnectionBuilder;

		private bool _disposed = false;
		private HubConnection _connection;

		protected SearchHubClientBase(IHubConnectionBuilder hubConnectionBuilder)
		{
			_hubConnectionBuilder = hubConnectionBuilder;
		}

		protected abstract Uri HubUrl { get; }

		public async Task InitializeHubConnection()
		{
			await GetConnection().ConfigureAwait(false);
		}

		public abstract Task SearchOnCallerAsync(SearchResponse message);

		public abstract Task SearchCalledOnAllAsync(SearchRequest message);

		public abstract Task SearchTreeRefreshedOnCallerAsync();

		public async Task CallRefreshSearchTreeOnHub()
		{
			var connection = await GetConnection().ConfigureAwait(false);

			await connection
				.SendAsync("RefreshSearchTree")
				.ConfigureAwait(false);
		}

		public async Task CallTriggerOnHub()
		{
			var connection = await GetConnection().ConfigureAwait(false);

			await connection
				.SendAsync("Trigger")
				.ConfigureAwait(false);
		}

		public async Task CallSearchOnHub(SearchRequest message)
		{
			var connection = await GetConnection().ConfigureAwait(false);

			await connection
				.SendAsync("Search", message)
				.ConfigureAwait(false);
		}

		private async Task<HubConnection> GetConnection()
		{
			if (_connection == null)
			{
				_connection = await CreateHubConnectionAsync().ConfigureAwait(false);
			}
		
			return _connection;
		}

		private async Task<HubConnection> CreateHubConnectionAsync()
		{
			var connection = _hubConnectionBuilder
				.WithUrl(HubUrl)
				.WithAutomaticReconnect(new[] { TimeSpan.Zero, TimeSpan.Zero, TimeSpan.FromSeconds(10) })
				.Build();

			BindClientMethods(ref connection);

			await connection
				.StartAsync()
				.ConfigureAwait(false);

			return connection;
		}

		private HubConnection BindClientMethods(ref HubConnection connection)
		{
			connection.On<SearchRequest>("SearchCalledOnAllAsync", message => SearchCalledOnAllAsync(message));
			connection.On("SearchTreeRefreshedOnCallerAsync", () => SearchTreeRefreshedOnCallerAsync());
			connection.On<SearchResponse>("SearchOnCallerAsync", message => SearchOnCallerAsync(message));

			return connection;
		}

		public void Dispose()
		{
			Dispose(true);
			GC.SuppressFinalize(this);
		}
		
		protected virtual void Dispose(bool disposing)
		{
			if (!_disposed)
			{
				if (disposing)
				{
					if (!(_connection is null))
					{
						var result = _connection
							.DisposeAsync()
							.Wait(1000);

						if (!result)
						{
							throw new TimeoutException($"Timeout of one second reached while closing hub connection for {HubUrl}!");
						}
					}
				}

				_disposed = true;
			}
		}
		
		~SearchHubClientBase()
		{
			Dispose(false);
		}
	}

	public abstract class SearchHubBase : Hub<ISearchHubClientBase>
	{
		protected Task SendSearchCalledOnAllAsync(SearchRequest message)
		{
			return Clients.All.SearchCalledOnAllAsync(message);
		}

		public async Task RefreshSearchTree()
		{
			await HandleRefreshSearchTreeAsync().ConfigureAwait(false);

			await Clients
				.Caller
				.SearchTreeRefreshedOnCallerAsync()
				.ConfigureAwait(false);
		}

		protected abstract Task HandleRefreshSearchTreeAsync();
		
		public async Task Trigger()
		{
			await HandleTriggerAsync().ConfigureAwait(false);
		}

		protected abstract Task HandleTriggerAsync();
		public async Task Search(SearchRequest message)
		{
			var result = await HandleSearchAsync(message).ConfigureAwait(false);

			await Clients
				.Caller
				.SearchOnCallerAsync(result)
				.ConfigureAwait(false);
		}

		protected abstract Task<SearchResponse> HandleSearchAsync(SearchRequest message);
		
	}

	public class SearchRequest
	{
		public int result_per_page { get; set; }
		public int page_number { get; set; }
		public string query { get; set; }
	}

	public class SearchResponse
	{
		public Result result { get; set; }
	}

	public class Result
	{
		public string snippets { get; set; }
		public string title { get; set; }
		public string url { get; set; }
	}
}